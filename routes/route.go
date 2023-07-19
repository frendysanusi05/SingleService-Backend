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

	admin := r.Group("")
	admin.Use(middlewares.CheckAdmin())
	{
		/******* PERUSAHAAN ********/
		admin.POST("/perusahaan", controllers.CreatePerusahaan)
		admin.PUT("/perusahaan/:id", controllers.UpdatePerusahaan)
		admin.DELETE("/perusahaan/:id", controllers.DeletePerusahaan)
		/******* BARANG ********/
		admin.POST("/barang", controllers.CreateBarang)
		admin.PUT("/barang/:id", controllers.UpdateBarang)
		admin.DELETE("/barang/:id", controllers.DeleteBarang)
	}
	/******* PERUSAHAAN ********/
	r.GET("/perusahaan", controllers.GetPerusahaan)
	r.GET("/perusahaan/:id", controllers.GetPerusahaanByID)
	/******* BARANG ********/
	r.GET("/barang", controllers.GetBarang)
	r.GET("/barang/:id", controllers.GetBarangByID)

	r.Run(":8080")
}