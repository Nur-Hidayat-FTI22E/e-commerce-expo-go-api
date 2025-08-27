package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/syrlramadhan/e-commerce-expo-go-api/config"
	"github.com/syrlramadhan/e-commerce-expo-go-api/controller"
	"github.com/syrlramadhan/e-commerce-expo-go-api/helper"
	"github.com/syrlramadhan/e-commerce-expo-go-api/repository"
	"github.com/syrlramadhan/e-commerce-expo-go-api/service"
)

func migrateDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := config.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Could not create database driver: %v", err)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	// Gabungkan direktori kerja dengan nama folder migrasi
	migrationsPath := "file://" + filepath.Join(currentDir, "migrations")

	m, err := migrate.NewWithDatabaseInstance(
		// "file://migrations",
		migrationsPath,
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("Migration failed to initialize: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while migrating: %v", err)
	}

	fmt.Println("Migration successful!")
}

func main() {
	migrateDB()
	errEnv := godotenv.Load()
	if errEnv != nil {
		helper.WriteJSONError(nil, http.StatusInternalServerError, "Failed to load environment variables: "+errEnv.Error())
		return
	}
	appPort := os.Getenv("APP_PORT")
	fmt.Println("api running on http://localhost:" + appPort)

	db, err := config.ConnectToDatabase()
	if err != nil {
		helper.WriteJSONError(nil, http.StatusInternalServerError, "Failed to connect to database: "+err.Error())
		return
	}

	router := httprouter.New()

	userRepo := repository.NewUserRepositoryImpl()
	userService := service.NewUserServiceImpl(userRepo, db)
	userController := controller.NewUserControllerImpl(userService)

	router.POST("/api/v1/user", userController.CreateUser)
	router.POST("/api/v1/user/login", userController.LoginUser)

	server := &http.Server{
		Addr:    ":" + appPort,
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		helper.WriteJSONError(nil, http.StatusInternalServerError, "Failed to start server: "+err.Error())
		return
	}
}
