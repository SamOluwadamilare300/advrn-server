package storage

import (
	"advrn-server/models/models"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectToDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not loaded, using system environment variables")
	}

	// Try URI format first
	dsn := os.Getenv("DB_CONNECTION_STRING")
	if dsn == "" {
		// Build DSN from individual components if URI not available
		dsn = buildDSNFromComponents()
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	log.Println("Database connection established")
	return db
}

func buildDSNFromComponents() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "require"
	}
	options := os.Getenv("DB_OPTIONS")

	// Build the DSN string
	dsn := []string{
		"host=" + host,
		"user=" + user,
		"password=" + password,
		"dbname=" + dbname,
		"port=" + port,
		"sslmode=" + sslmode,
	}
	if options != "" {
		dsn = append(dsn, "options="+options)
	}

	return strings.Join(dsn, " ")
}

func performMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Conversation{},
		&models.Message{},
		&models.User{},
		&models.Property{},
		&models.Review{},
		&models.Apartment{},
	)
	if err != nil {
		log.Fatalf("Failed to perform migrations: %v", err)
	}
	log.Println("Migrations completed successfully")
}

func InitializeDB() *gorm.DB {
	db := connectToDB()
	performMigrations(db)
	return db
}