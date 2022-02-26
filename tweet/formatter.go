package tweet

type tweetFormat struct {
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

type tweetsFormat struct {
	Id        uint   `json:"id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	User      userFormat
}

type userFormat struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func TweetFormatter(tweet Tweet) tweetFormat {
	newStatus := tweetFormat{Status: tweet.Status, CreatedAt: tweet.CreatedAt.String()}

	return newStatus
}

func TweetsFormatter(tweets []Tweet) []tweetsFormat {
	res := []tweetsFormat{}
	for _, tweet := range tweets {
		user := userFormat{
			Id:       tweet.User.ID,
			Name:     tweet.User.Name,
			Username: tweet.User.Username,
		}
		format := tweetsFormat{
			Id:        tweet.ID,
			Status:    tweet.Status,
			CreatedAt: tweet.CreatedAt.String(),
			User:      user,
		}
		res = append(res, format)
	}

	return res
}
