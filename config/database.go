package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Ganti dengan konfigurasi database Anda
	dsn := "root:@tcp(127.0.0.1:3306)/url_shortener_db?charset=utf8mb4&parseTime=True&loc=Local"
	
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database!", err)
	}

	fmt.Println("Koneksi ke database berhasil!")
	DB = database
}