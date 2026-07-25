package repository

type User struct {
	Username string
	Password string
}

func GetUser() User {
	return User{
		Username: "admin",
		Password: "12345",
	}
}