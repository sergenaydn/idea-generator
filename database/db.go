package database

import (
	"idea-generator/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a global variable to hold the database connection.
var DB *gorm.DB

// Init initializes the database connection and sets up the required configurations.
func Init() {
	// Define the data source name (DSN) for the PostgreSQL database.
	// It contains the necessary connection details such as username, password, host, and database name.
	dsn := "postgres://root:root@localhost:5432/ideas"

	// Attempt to establish a connection to the PostgreSQL database using the specified DSN.
	// The gorm.Open function takes the database driver (postgres) and the DSN as arguments.
	// It returns a *gorm.DB instance representing the database connection.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// If there's an error while connecting to the database, the application panics (exits with an error).
		panic(err)
	}

	// AutoMigrate is used to automatically create the necessary database tables for the given model(s).
	// In this case, it ensures that the "Idea" model's table exists in the database.
	db.AutoMigrate(&models.Idea{})

	// Store the database connection in the global variable DB so that it can be used throughout the application.
	DB = db
}
