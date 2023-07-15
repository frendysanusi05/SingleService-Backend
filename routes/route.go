package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/frendysanusi05/Seleksi-Asisten-Laboratorium-Programming-SingleService/controllers"
	"github.com/frendysanusi05/Seleksi-Asisten-Laboratorium-Programming-SingleService/middlewares"
)

func Route() {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	r.POST("/login", controllers.Login)
	r.GET("/self", controllers.Self)

	r.Run(":8080")
}