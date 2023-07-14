package controllers

import (
	"net/http"
  	"github.com/gin-gonic/gin"
	"github.com/frendysanusi05/Seleksi-Asisten-Laboratorium-Programming-SingleService/models"
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

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	// set cookie
	cookie, err := c.Cookie("gin_cookie")

	if err != nil {
		cookie = token
		c.SetCookie("gin_cookie", cookie, 3600, "/", "localhost", false, true)
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
