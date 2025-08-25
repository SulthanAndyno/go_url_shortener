package controllers

import (
	"net/http"
		"github.com/SulthanAndyno/url-shortener/config"
		"github.com/SulthanAndyno/url-shortener/models"
		"github.com/SulthanAndyno/url-shortener/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShortenRequest struct {
	OriginalURL string `json:"original_url" binding:"required,url"`
}

// POST /shorten
func ShortenURL(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL tidak valid"})
		return
	}

	shortCode := utils.GenerateShortCode()

	// Pastikan shortCode unik (opsional, untuk aplikasi skala besar)
	// Untuk proyek sederhana ini, kita asumsikan tidak ada kolisi

	url := models.Url{
		OriginalURL: req.OriginalURL,
		ShortCode:   shortCode,
	}

	result := config.DB.Create(&url)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan URL"})
		return
	}
	
	// Ganti "localhost:8080" dengan domain Anda jika di-deploy
	shortURL := "http://localhost:8080/" + shortCode

	c.JSON(http.StatusCreated, gin.H{
		"message":   "URL berhasil dipendekkan!",
		"short_url": shortURL,
	})
}

// GET /:short_code
func RedirectURL(c *gin.Context) {
	shortCode := c.Param("short_code")
	var url models.Url

	result := config.DB.Where("short_code = ?", shortCode).First(&url)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short code tidak ditemukan"})
		return
	}

	// Tambahkan click_count +1 secara atomik
	config.DB.Model(&url).Update("click_count", gorm.Expr("click_count + 1"))

	c.Redirect(http.StatusFound, url.OriginalURL)
}

// GET /stats/:short_code
func GetURLStats(c *gin.Context) {
	shortCode := c.Param("short_code")
	var url models.Url

	result := config.DB.Where("short_code = ?", shortCode).First(&url)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short code tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"original_url": url.OriginalURL,
		"short_code":   url.ShortCode,
		"click_count":  url.ClickCount,
		"created_at":   url.CreatedAt,
	})
}