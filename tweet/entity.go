package tweet

import (
	"workshop-be/user"

	"gorm.io/gorm"
)

type Tweet struct {
	gorm.Model
	Status string
	UserID uint
	User   user.User
}

type InputAddTweet struct {
	Status string `binding:"required"`
	UserID uint
}

type InputUpdateTweet struct {
	Status  string `binding:"required"`
	UserID  uint
	TweetID uint
}

type InputUriTweet struct {
	Id int `uri:"id" binding:"required"`
}
