package tweet

import "errors"

type Service interface {
	CreateTweet(input InputAddTweet) (Tweet, error)
	UpdateTweet(input InputUpdateTweet) (Tweet, error)
	GetTweetsByUserId(userId uint) ([]Tweet, error)
	GetTweets() ([]Tweet, error)
	GetTweetById(id uint) (Tweet, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateTweet(input InputAddTweet) (Tweet, error) {
	tweet := Tweet{
		Status: input.Status,
		UserID: input.UserID,
	}

	tweet, err := s.repository.Save(tweet)
	if err != nil {
		return tweet, errors.New("gagal membuat status")
	}

	return tweet, nil
}

func (s *service) UpdateTweet(input InputUpdateTweet) (Tweet, error) {
	tweet, err := s.repository.GetById(input.TweetID)
	if err != nil {
		return tweet, err
	}

	if tweet.ID == 0 {
		return tweet, errors.New("tweet tidak ditemukan")
	}

	if tweet.UserID != input.UserID {
		return tweet, errors.New("tidak memiliki akses")
	}

	tweet.Status = input.Status

	newTweet, err := s.repository.Update(tweet)
	if err != nil {
		return tweet, err
	}

	return newTweet, nil
}

func (s *service) GetTweetsByUserId(userId uint) ([]Tweet, error) {
	var tweets []Tweet
	tweets, err := s.repository.GetByUserId(userId)
	if err != nil {
		return tweets, err
	}

	return tweets, nil
}

func (s *service) GetTweets() ([]Tweet, error) {
	var tweets []Tweet
	tweets, err := s.repository.Get()
	if err != nil {
		return tweets, err
	}

	return tweets, nil
}

func (s *service) GetTweetById(id uint) (Tweet, error) {
	var tweet Tweet
	tweet, err := s.repository.GetById(id)
	if err != nil {
		return tweet, err
	}

	return tweet, nil
}
