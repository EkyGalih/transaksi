package db

import (
	"fmt"

	"github.com/EkyGalih/lkpd-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	conf := config.GetConfig()

	// String koneksi untuk PostgreSQL
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Makassar",
		conf.DB_HOST, conf.DB_PORT, conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_NAME,
	)

	// Membuka koneksi dengan PostgreSQL
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Error opening database: %s", err.Error())) // Menampilkan error yang lebih jelas
	}
}

func CreateCon() *gorm.DB {
	return db
}
