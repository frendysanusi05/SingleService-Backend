package controllers

import (
	"fmt"
	"context"
	"net/http"
	"io/ioutil"

  	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/frendysanusi05/Seleksi-Asisten-Laboratorium-Programming-SingleService/models"
	tokens "github.com/frendysanusi05/Seleksi-Asisten-Laboratorium-Programming-SingleService/utils/token"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.GetUser()

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	// set cookie
	cookie, err := c.Cookie("cookie")

	if err != nil {
		cookie = token
		c.SetCookie("cookie", cookie, 3600, "/", "localhost", false, true)
		// c.Writer.Header().Set("Authorization", cookie)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Login success",
		"data": gin.H{
			"user": gin.H{
				"username": u.Username,
				// "name": u.Name,
			},
			"token": token,
		},
	})
}

func Self(c *gin.Context) {
	tokenString := tokens.ExtractToken(c)
	token, err := tokens.ParseToken(tokenString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parsing token error"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println("claims: ", claims)
	if !ok || !tokens.TokenValid(c) {
		fmt.Println("ok: ", ok)
		fmt.Println("tokenValid: ", tokens.TokenValid(c))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ctx := context.WithValue(context.Background(), "claims", claims)
	req := c.Request.WithContext(ctx)

	userInfo := req.Context().Value("claims").(jwt.MapClaims)

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
				// "name": userInfo["name"],
			},
		},
	})
}
