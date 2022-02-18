package user

type UserFormat struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func FormatterUser(user User, token string) UserFormat {
	newUser := UserFormat{
		Id:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Token:    token,
	}

	return newUser
}
