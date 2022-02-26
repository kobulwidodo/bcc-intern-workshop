package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"workshop-be/auth"
	"workshop-be/config"
	"workshop-be/handler"
	"workshop-be/helper"
	"workshop-be/tweet"
	"workshop-be/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
	tweetRepository := tweet.NewRepository(db)

	// setup Service
	userService := user.NewService(userRepository)
	authService := auth.NewJwtService()
	tweetService := tweet.NewService(tweetRepository)

	// setup handler
	userHandler := handler.NewUserHandler(userService, authService)
	tweetHandler := handler.NewTweetHandler(tweetService)

	api := r.Group("/api")
	{
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)

		api.POST("/tweet", authMiddleware(authService, userService), tweetHandler.CreateTweet)
		api.PUT("/tweet/:id", authMiddleware(authService, userService), tweetHandler.UpdateTweet)
		api.GET("/tweet/user/:id", tweetHandler.GetTweetsByUserId)
		api.GET("/tweet", tweetHandler.GetAllTweets)
		api.GET("/tweet/:id", tweetHandler.GetTweetById)
		api.DELETE("/tweet/:id", authMiddleware(authService, userService), tweetHandler.DeleteTweet)
	}

	r.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Not a Bearer token", http.StatusUnauthorized, "gagal", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ApiResponse("Token Invalid", http.StatusUnauthorized, "gagal", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "gagal", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := uint(claim["user_id"].(float64))
		var user user.User
		user, err = userService.GetById(userId)
		if err != nil {
			response := helper.ApiResponse("Failed to get user", http.StatusUnauthorized, "gagal", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("userLoggedin", user)
	}
}
