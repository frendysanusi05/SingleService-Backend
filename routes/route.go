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

	/******* PERUSAHAAN ********/
	r.GET("/perusahaan", controllers.GetPerusahaan)
	r.POST("/perusahaan", controllers.CreatePerusahaan)
	r.GET("/perusahaan/:id", controllers.GetPerusahaanByID)
	r.PUT("/perusahaan/:id", controllers.UpdatePerusahaan)
	r.DELETE("/perusahaan/:id", controllers.DeletePerusahaan)

	/******* BARANG ********/
	r.GET("/barang", controllers.GetBarang)
	r.POST("/barang", controllers.CreateBarang)
	r.GET("/barang/:id", controllers.GetBarangByID)
	r.PUT("/barang/:id", controllers.UpdateBarang)
	r.DELETE("/barang/:id", controllers.DeleteBarang)

	r.Run(":8080")
}