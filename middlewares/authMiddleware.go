package middlewares

import (	
	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"

	tokens "single-service/utils/token"
	"single-service/utils"
)

func CheckAdmin() gin.HandlerFunc {
	return func (c *gin.Context) {
		tokenString := tokens.ExtractToken(c)
		token, err := tokens.ParseToken(tokenString)

		if err != nil {
			utils.MessageBadRequest(c, "An error occured")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !tokens.TokenValid(c) {
			utils.MessageUnauthorized(c, "Unauthorized")
			return
		}

		if role, ok := claims["role"].(string); !ok || role != "admin" {
			utils.MessageUnauthorized(c, "Unauthorized")
			return
		}

		c.Next()
	}
}
