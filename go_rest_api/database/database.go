package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/123-zuleyha/go_rest_api/config"
	"github.com/123-zuleyha/go_rest_api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database instance
type Dbinstance struct {
	Db *gorm.DB // GORM veritabanı nesnesi
}

var DB Dbinstance // Veritabanı örneği

// Connect function database
func Connect() {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Fatalf("Error parsing port: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.Config("DB_HOST"),
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
		port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database")
	log.Println("Running migrations...")
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	DB = Dbinstance{
		Db: db,
	}
}
