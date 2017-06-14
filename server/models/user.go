package models

type User struct {
    Id int `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func (user *User) GetUsername() string {
    return user.Username
}

func (user *User) GetPassword() string {
    return user.Password
}
