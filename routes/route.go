package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/frendysanusi05/Seleksi-Asisten-Laboratorium-Programming-SingleService/controllers"
	"github.com/frendysanusi05/Seleksi-Asisten-Laboratorium-Programming-SingleService/middlewares"
	// tokens "github.com/frendysanusi05/Seleksi-Asisten-Laboratorium-Programming-SingleService/utils/token"
)

func Route() {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	// r.Use(tokens.SetCookies())
	r.POST("/login", controllers.Login)

	r.Run(":8080")
}