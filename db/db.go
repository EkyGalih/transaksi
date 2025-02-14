package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/EkyGalih/transaksi/config"
	"github.com/EkyGalih/transaksi/entities"
	_ "github.com/lib/pq" // Driver PostgreSQL untuk sql.Open()
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	conf := config.GetConfig()

	// Koneksi awal ke PostgreSQL tanpa memilih database tertentu
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s sslmode=disable TimeZone=Asia/Makassar",
		conf.DB_HOST, conf.DB_PORT, conf.DB_USERNAME, conf.DB_PASSWORD,
	)

	// Gunakan sql.Open untuk mengecek database
	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}
	defer sqlDB.Close()

	// Cek apakah database sudah ada
	var exists bool
	query := fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname='%s'", conf.DB_NAME)
	err = sqlDB.QueryRow(query).Scan(&exists)

	if err != nil {
		// Jika database belum ada, buat database
		if err == sql.ErrNoRows {
			log.Printf("Database %s tidak ditemukan, membuat database...", conf.DB_NAME)
			createDBQuery := fmt.Sprintf("CREATE DATABASE %s", conf.DB_NAME)
			_, err = sqlDB.Exec(createDBQuery)
			if err != nil {
				log.Fatal("Gagal membuat database:", err)
			}
			log.Println("Database berhasil dibuat.")
		} else {
			log.Fatal("Error checking database:", err)
		}
	} else {
		log.Println("Database sudah ada, melanjutkan koneksi...")
	}

	// Setelah database ada, koneksi ulang dengan database yang benar
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Makassar",
		conf.DB_HOST, conf.DB_PORT, conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_NAME,
	)

	// Membuka koneksi dengan PostgreSQL menggunakan GORM
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Cek apakah tabel transactions sudah ada sebelum AutoMigrate
	if !db.Migrator().HasTable(&entities.Transaction{}) {
		log.Println("Table 'transactions' not found, running migration...")
		err = db.AutoMigrate(&entities.Transaction{})
		if err != nil {
			log.Fatal("Migration failed:", err)
		}
		log.Println("Migration completed successfully.")
	} else {
		log.Println("Table 'transactions' already exists, skipping migration.")
	}
}

// Fungsi untuk mendapatkan koneksi database
func CreateCon() *gorm.DB {
	return db
}
