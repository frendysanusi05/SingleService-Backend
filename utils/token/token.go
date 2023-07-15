package token

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string, password string) (string, error) {
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "",err
	}

	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func TokenValid(c *gin.Context) bool {
	tokenString := ExtractToken(c)
	_, err := ParseToken(tokenString)

	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	return true
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
}

func ExtractToken(c *gin.Context) string {
	// cookie, err := c.Request.Cookie("cookie")
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return ""
	// }

	// return cookie.Value
	return c.GetHeader("Authorization")
}
