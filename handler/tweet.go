package handler

import (
	"net/http"
	"workshop-be/helper"
	"workshop-be/tweet"
	"workshop-be/user"

	"github.com/gin-gonic/gin"
)

type tweetHandler struct {
	tweetService tweet.Service
}

func NewTweetHandler(tweetService tweet.Service) *tweetHandler {
	return &tweetHandler{tweetService: tweetService}
}

func (h *tweetHandler) CreateTweet(c *gin.Context) {
	var input tweet.InputAddTweet
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.ApiResponse("Harap mengisi semua input", http.StatusUnprocessableEntity, "error", helper.FormatBindError(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLoggedin := c.MustGet("userLoggedin").(user.User)
	input.UserID = userLoggedin.ID

	newTweet, err := h.tweetService.CreateTweet(input)
	if err != nil {
		response := helper.ApiResponse(err.Error(), http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.ApiResponse("Berhasil membuat tweet baru", http.StatusOK, "success", tweet.TweetFormatter(newTweet))
	c.JSON(http.StatusOK, response)
}

func (h *tweetHandler) UpdateTweet(c *gin.Context) {
	var inputId tweet.InputUriTweet
	if err := c.ShouldBindUri(&inputId); err != nil {
		response := helper.ApiResponse("ID tidak tersedia", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var input tweet.InputUpdateTweet
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.ApiResponse("Harap isi semua input", http.StatusUnprocessableEntity, "error", helper.FormatBindError(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLoggedin := c.MustGet("userLoggedin").(user.User)
	input.UserID = userLoggedin.ID
	input.TweetID = uint(inputId.Id)

	updatedtweet, err := h.tweetService.UpdateTweet(input)
	if err != nil {
		response := helper.ApiResponse(err.Error(), http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.ApiResponse("Sukses mengubah tweet", http.StatusOK, "success", tweet.TweetFormatter(updatedtweet))
	c.JSON(http.StatusOK, response)
}

func (h *tweetHandler) GetTweetsByUserId(c *gin.Context) {
	var input tweet.InputUriTweet
	if err := c.ShouldBindUri(&input); err != nil {
		response := helper.ApiResponse("User id tidak ditemukan", http.StatusUnprocessableEntity, "error", helper.FormatBindError(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var tweets []tweet.Tweet
	tweets, err := h.tweetService.GetTweetsByUserId(uint(input.Id))
	if err != nil {
		response := helper.ApiResponse(err.Error(), http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.ApiResponse("Berhasil mendapatkan data", http.StatusOK, "success", tweet.TweetsFormatter(tweets))
	c.JSON(http.StatusOK, response)
}

func (h *tweetHandler) GetAllTweets(c *gin.Context) {
	var tweets []tweet.Tweet
	tweets, err := h.tweetService.GetTweets()
	if err != nil {
		response := helper.ApiResponse("Gagal mendapatkan data", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.ApiResponse("Berhasil mendapatkan data", http.StatusOK, "success", tweet.TweetsFormatter(tweets))
	c.JSON(http.StatusOK, response)
}

func (h *tweetHandler) GetTweetById(c *gin.Context) {
	var input tweet.InputUriTweet
	if err := c.ShouldBindUri(&input); err != nil {
		response := helper.ApiResponse("Id tidak ditemukan", http.StatusUnprocessableEntity, "error", helper.FormatBindError(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	getTweet, err := h.tweetService.GetTweetById(uint(input.Id))
	if err != nil {
		response := helper.ApiResponse("Gagal mendapatkan data", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.ApiResponse("Berhasil mendapatkan data", http.StatusOK, "success", tweet.TweetFormatter(getTweet))
	c.JSON(http.StatusOK, response)
}
