package routes

import (
	"github.com/gin-gonic/gin"
	"single-service/controllers"
	"single-service/middlewares"
)

func Route() {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	r.POST("/login", controllers.Login)
	r.GET("/self", controllers.Self)

	r.Run(":8080")
}