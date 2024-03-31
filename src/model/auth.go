package model

import "bs-server/src/bo"

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Icon     string `json:"icon"`
}

func CheckAuth(username, password string) bool {
	var user User

	res := db.Where(User{
		Username: username,
		Password: password,
	}).First(&user)

	return res.Error == nil
}

func Register(user *bo.RegisterBO) bool {
	res := db.Create(user)

	return res.Error != nil
}
