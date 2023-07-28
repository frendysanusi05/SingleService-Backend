package controllers

import (
	"net/http"
	"strconv"
	"errors"

	"github.com/gin-gonic/gin"

	"single-service/models"
	"single-service/databases"
	"single-service/utils"
)

func GetBarang(c *gin.Context) {
	query := c.Query("q")

	var barang []models.Barang
	DB, _ := databases.ConnectDatabase()

	if query != "" {
		DB = DB.Where("nama LIKE ? OR kode LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	if err := DB.Find(&barang).Error; err != nil {
		utils.MessageBadRequest(c, "An error occured")
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

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "GET barang success",
		"data": barangData,
	})
}

func GetBarangByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var barang models.Barang
	DB, _ := databases.ConnectDatabase()

	if err := DB.Where("id = ?", id).First(&barang).Error; err != nil {
		utils.MessageBadRequest(c, "An error occured")
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "GET barang success",
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
		utils.MessageBadRequest(c, "Each inputs must have a value")
		return
	}

	err := validateBarang(c, barang)
	if err != nil {
		utils.MessageBadRequest(c, err.Error())
		return
	}

	if barang.ID == "" {
		var lastBarang models.Barang
		res := DB.Order("CONVERT(id, SIGNED) DESC").Select("id").First(&lastBarang)
		if res.Error == nil {
			lastIDInt, _ := strconv.Atoi(lastBarang.ID)
			lastIDInt = lastIDInt + 1
			barang.ID = strconv.Itoa(lastIDInt)
		} else {
			barang.ID = "1"
		}
	}

	if err := DB.Create(&barang).Error; err != nil {
		utils.MessageInternalError(c, "An error occured")
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "POST barang success",
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

func UpdateBarang(c *gin.Context) {
	id := c.Params.ByName("id")
	var barang models.Barang
	DB, _ := databases.ConnectDatabase()

	if err := DB.Where("id = ?", id).First(&barang).Error; err != nil {
		utils.MessageBadRequest(c, "An error occured")
		return
	}
	
	if err := c.ShouldBindJSON(&barang); err != nil {
		utils.MessageBadRequest(c, "Each inputs must have a value")
		return
	}

	err := validateBarang(c, barang)
	if err != nil {
		utils.MessageBadRequest(c, err.Error());
		return
	}

	DB.Save(&barang)
	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Update barang success",
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

func DeleteBarang(c *gin.Context) {
	id := c.Params.ByName("id")
	var barang models.Barang
	DB, _ := databases.ConnectDatabase()

	if err := DB.Where("id = ?", id).First(&barang).Error; err != nil {
		utils.MessageInternalError(c, "An error occured")
		return
	}

	deletedBarang := barang

	if err := DB.Delete(&barang).Error; err != nil {
		utils.MessageInternalError(c, "An error occured")
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Delete barang success",
		"data": gin.H{
			"id": deletedBarang.ID,
			"nama": deletedBarang.Nama,
			"harga": deletedBarang.Harga,
			"stok": deletedBarang.Stok,
			"kode": deletedBarang.Kode,
			"perusahaan_id": deletedBarang.PerusahaanID,
		},
	})	
}

/******** ADDITIONAL FUNCTION *********/
func validateBarang(c *gin.Context, barang models.Barang) error {
	if !validateHarga(barang) {
		return errors.New("Harga must be greater than 0")
	}

	if !validateStok(barang) {
		return errors.New("Stok cannot be negative")
	}

	if !validateKodeBarang(barang) {
		return errors.New("Kode must be unique")
	}
	return nil
}

func validateHarga(b models.Barang) bool {
	if b.Harga > 0 {
		return true
	}
	return false
}

func validateStok(b models.Barang) bool {
	if b.Stok >= 0 {
		return true
	}
	return false
}

func validateKodeBarang(b models.Barang) bool {
	DB, _ := databases.ConnectDatabase()
	var barangs []models.Barang
	if err := DB.Where("kode = ?", b.Kode).Find(&barangs).Error; err != nil {
		return false
	}
	
	for _, barang := range barangs {
		if barang.Kode == b.Kode && barang.ID != b.ID {
			return false
		}
	}
	return true
}
