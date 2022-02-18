package main

import (
	"fmt"
	"log"
	"workshop-be/auth"
	"workshop-be/config"
	"workshop-be/handler"
	"workshop-be/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.InitDB()
	if err != nil {
		fmt.Println(err)
		panic("Cant connect to database")
	}

	// setup Repository
	userRepository := user.NewRepository(db)

	// setup Service
	userService := user.NewService(userRepository)
	authService := auth.NewJwtService()

	userHandler := handler.NewUserHandler(userService, authService)

	api := r.Group("/api")
	{
		api.POST("/register", userHandler.Register)
	}

	r.Run()
}
