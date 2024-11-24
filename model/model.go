package model

type User struct {
	Username string
	Password string
	Email string
}

func GetUserByUsername(username string) (User, error) {
	var u User
	return u, nil	
}