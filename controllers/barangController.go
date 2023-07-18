package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"single-service/models"
	"single-service/databases"
)

func GetBarang(c *gin.Context) {
	query := c.Query("q")

	var barang []models.Barang
	DB, _ := databases.ConnectDatabase()

	if query != "" {
		DB = DB.Where("nama LIKE ? OR kode LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	if err := DB.Find(&barang).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var barangData []gin.H
	for _, b := range barang {
		barangData = append(barangData, gin.H{
			"id": b.ID,
			"nama": b.Nama,
			"harga": b.Harga,
			"stok": b.Stok,
			"kode": b.Kode,
			"perusahaan_id": b.PerusahaanID,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "GET Barang success",
		"data": barangData,
	})
}

func GetBarangByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var barang models.Barang
	DB, _ := databases.ConnectDatabase()

	if err := DB.Where("id = ?", id).First(&barang).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "GET Barang success",
		"data": gin.H{
			"id": barang.ID,
			"nama": barang.Nama,
			"harga": barang.Harga,
			"stok": barang.Stok,
			"kode": barang.Kode,
			"perusahaan_id": barang.PerusahaanID,
		},
	})
}

func CreateBarang(c *gin.Context) {
	var barang models.Barang
	DB, _ := databases.ConnectDatabase()

	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if barang.ID == "" {
		var lastBarang models.Barang
		res := DB.Order("id DESC").Select("id").First(&lastBarang)
		if res.Error == nil {
			lastIDInt, _ := strconv.Atoi(lastBarang.ID)
			lastIDInt = lastIDInt + 1
			barang.ID = strconv.Itoa(lastIDInt)
		} else {
			barang.ID = "1"
		}
	}

	if err := DB.Create(&barang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create barang"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "POST Barang success",
		"data": gin.H{
			"id": barang.ID,
			"nama": barang.Nama,
			"harga": barang.Harga,
			"stok": barang.Stok,
			"kode": barang.Kode,
			"perusahaan_id": barang.PerusahaanID,
		},
	})
}
