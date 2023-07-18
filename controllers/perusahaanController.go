package controllers

import (
	"strconv"
	"strings"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"

	"single-service/models"
	"single-service/databases"
)

func GetPerusahaan(c *gin.Context) {
	query := c.Query("q")

	var perusahaan []models.Perusahaan
	DB, _ := databases.ConnectDatabase()

	if query != "" {
		DB = DB.Where("nama LIKE ? OR kode LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	if err := DB.Find(&perusahaan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
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

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "GET Perusahaan success",
		"data": perusahaanData,
	})
}

func GetPerusahaanByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var perusahaan models.Perusahaan
	DB, _ := databases.ConnectDatabase()

	if err := DB.Where("id = ?", id).First(&perusahaan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "GET Perusahaan success",
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	if !ValidateKode(perusahaan) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kode harus berupa 3 huruf kapital"})
		return
	}

	if err := DB.Create(&perusahaan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create perusahaan"})
		return
	}

	perusahaanData := gin.H{
		"id": perusahaan.ID,
		"nama": perusahaan.Nama,
		"alamat": perusahaan.Alamat,
		"no_telp": perusahaan.NoTelp,
		"kode": perusahaan.Kode,
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "POST Perusahaan success",
		"data": perusahaanData,
	})
}

/********* ADDITIONAL FUNCTION *********/
func ValidateKode(p models.Perusahaan) bool {
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
