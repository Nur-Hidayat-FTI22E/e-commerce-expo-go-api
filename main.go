package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/syrlramadhan/e-commerce-expo-go-api/config"
	"github.com/syrlramadhan/e-commerce-expo-go-api/controller"
	"github.com/syrlramadhan/e-commerce-expo-go-api/helper"
	"github.com/syrlramadhan/e-commerce-expo-go-api/repository"
	"github.com/syrlramadhan/e-commerce-expo-go-api/service"
)

func main() {
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
