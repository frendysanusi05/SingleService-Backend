package controllers

import (
	"strconv"
	"strings"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"

	"single-service/models"
	"single-service/databases"
	"single-service/utils"
)

type PerusahaanInput struct {
	Nama 	string `json:"nama" binding:"required"`
	Alamat 	string `json:"alamat" binding:"required"`
	NoTelp 	string `json:"no_telp" binding:"required"`
	Kode 	string `json:"kode" binding:"required"`
}

func GetPerusahaan(c *gin.Context) {
	query := c.Query("q")

	var perusahaan []models.Perusahaan
	DB, _ := databases.ConnectDatabase()

	if query != "" {
		DB = DB.Where("nama LIKE ? OR kode LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	if err := DB.Find(&perusahaan).Error; err != nil {
		utils.MessageBadRequest(c, "An error occured")
		return
	}

	var perusahaanData []gin.H
	for _, p := range perusahaan {
		perusahaanData = append(perusahaanData, gin.H{
			"id": p.ID,
			"nama": p.Nama,
			"alamat": p.Alamat,
			"no_telp": p.NoTelp,
			"kode": p.Kode,
		})
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Get perusahaan success",
		"data": perusahaanData,
	})
}

func GetPerusahaanByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var perusahaan models.Perusahaan
	DB, _ := databases.ConnectDatabase()

	if err := DB.Where("id = ?", id).First(&perusahaan).Error; err != nil {
		utils.MessageBadRequest(c, "An error occured")
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "GET perusahaan success",
		"data": gin.H{
			"id": perusahaan.ID,
			"nama": perusahaan.Nama,
			"alamat": perusahaan.Alamat,
			"no_telp": perusahaan.NoTelp,
			"kode": perusahaan.Kode,
		},
	})
}

func CreatePerusahaan(c *gin.Context) {
	var perusahaan models.Perusahaan
	DB, _ := databases.ConnectDatabase()

	if err := c.ShouldBindJSON(&perusahaan); err != nil {
		utils.MessageBadRequest(c, "Each inputs must have a value")
		return
	}

	if perusahaan.ID == "" {
		var lastPerusahaan models.Perusahaan
		res := DB.Order("id DESC").Select("id").First(&lastPerusahaan)
		if res.Error == nil {
			lastIDInt, _ := strconv.Atoi(lastPerusahaan.ID)
			lastIDInt = lastIDInt + 1
			perusahaan.ID = strconv.Itoa(lastIDInt)
		} else {
			perusahaan.ID = "1"
		}
	}

	if !validateKodePerusahaan(perusahaan) {
		utils.MessageInternalError(c, "Kode must contain exactly 3 capital letter")
		return
	}

	if err := DB.Create(&perusahaan).Error; err != nil {
		utils.MessageInternalError(c, "An error occured")
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Create perusahaan success",
		"data": gin.H{
			"id": perusahaan.ID,
			"nama": perusahaan.Nama,
			"alamat": perusahaan.Alamat,
			"no_telp": perusahaan.NoTelp,
			"kode": perusahaan.Kode,
		},
	})
}

func UpdatePerusahaan(c *gin.Context) {
	id := c.Params.ByName("id")
	var perusahaan models.Perusahaan
	DB, _ := databases.ConnectDatabase()

	if err := DB.Where("id = ?", id).First(&perusahaan).Error; err != nil {
		utils.MessageBadRequest(c, "An error occured")
		return
	}
	
	if err := c.ShouldBindJSON(&perusahaan); err != nil {
		utils.MessageBadRequest(c, "Each inputs must have a value")
		return
	}

	if !validateKodePerusahaan(perusahaan) {
		utils.MessageInternalError(c, "Kode must contain exactly 3 capital letter")
		return
	}

	DB.Save(&perusahaan)
	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Update perusahaan success",
		"data": gin.H{
			"id": perusahaan.ID,
			"nama": perusahaan.Nama,
			"alamat": perusahaan.Alamat,
			"no_telp": perusahaan.NoTelp,
			"kode": perusahaan.Kode,
		},
	})
}

func DeletePerusahaan(c *gin.Context) {
	id := c.Params.ByName("id")
	var perusahaan models.Perusahaan
	DB, _ := databases.ConnectDatabase()

	if err := DB.Where("id = ?", id).First(&perusahaan).Error; err != nil {
		utils.MessageInternalError(c, "An error occured")
		return
	}

	deletedPerusahaan := perusahaan

	if err := DB.Delete(&perusahaan).Error; err != nil {
		utils.MessageInternalError(c, "An error occured")
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "Delete barang success",
		"data": gin.H{
			"id": deletedPerusahaan.ID,
			"nama": deletedPerusahaan.Nama,
			"alamat": deletedPerusahaan.Alamat,
			"no_telp": deletedPerusahaan.NoTelp,
			"kode": deletedPerusahaan.Kode,
		},
	})	
}

/********* ADDITIONAL FUNCTION *********/
func validateKodePerusahaan(p models.Perusahaan) bool {
	kode := strings.TrimSpace(p.Kode)
	if len(kode) != 3 {
		return false
	}

	match, err := regexp.MatchString("^[A-Z]{3}$", kode)
	if err != nil {
		return false
	}

	return match
}
