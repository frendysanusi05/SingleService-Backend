package controllers

import (
	"fmt"
	"context"
	"net/http"
	"io/ioutil"

  	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	jwt "github.com/dgrijalva/jwt-go"

	tokens "single-service/utils/token"
	"single-service/models"
	"single-service/databases"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	DB, _ := databases.ConnectDatabase()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := LoginCheck(u.Username, u.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	err = DB.Select("name").Where("username = ?", u.Username).First(&u).Error
	if err != nil {
		return
	}

	// set cookie
	cookie, err := c.Cookie("cookie")

	if err != nil {
		cookie = token
		c.SetCookie("cookie", cookie, 3600, "/", "localhost", false, true)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Login success",
		"data": gin.H{
			"user": gin.H{
				"username": u.Username,
				"name": u.Name,
			},
			"token": token,
		},
	})
}

func Self(c *gin.Context) {
	DB, _ := databases.ConnectDatabase()

	tokenString := tokens.ExtractToken(c)
	token, err := tokens.ParseToken(tokenString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parsing token error"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !tokens.TokenValid(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ctx := context.WithValue(context.Background(), "claims", claims)
	req := c.Request.WithContext(ctx)

	userInfo := req.Context().Value("claims").(jwt.MapClaims)

	// get name
	u := models.User{}
	err = DB.Select("name").Where("username = ?", userInfo["username"]).First(&u).Error
	if err != nil {
		return
	}

	resp, err := http.Get("https://ohl-fe.vercel.app/self")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Request GET failed"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Request body failed"})
		return
	}

	fmt.Println("Body: ", string(body))

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Read self success",
		"data": gin.H{
			"user": gin.H{
				"username": userInfo["username"],
				"name": u.Name,
			},
		},
	})
}

/**** ADDITIONAL FUNCTIONS *****/
func LoginCheck(username string, password string) (string,error) {	
	var err error

	u := models.User{}
	DB, _ := databases.ConnectDatabase()
	err = DB.Model(models.User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("Password salah")
		return "", err
	}

	token,err := tokens.GenerateToken(username, password)

	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyPassword(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

