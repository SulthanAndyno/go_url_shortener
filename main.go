package main

import (
	"github.com/SulthanAndyno/url-shortener/config"
	"github.com/SulthanAndyno/url-shortener/models"
	"github.com/SulthanAndyno/url-shortener/routes"
)

func main() {
	// 1. Hubungkan ke Database
	config.ConnectDatabase()

	// 2. Lakukan migrasi tabel Url
	config.DB.AutoMigrate(&models.Url{})

	// 3. Setup router
	router := routes.SetupRouter()

	// 4. Jalankan server di port 8080
	router.Run(":8080")
}