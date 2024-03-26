package model

import "bs-server/src/bo"

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Icon     string `json:"icon"`
}

func CheckAuth(username, password string) bool {
	var auth Auth

	db.Select("id").Where(Auth{
		Username: username,
		Password: password,
	}).First(&auth)

	return auth.Username != ""
}

func Register(user *bo.RegisterBO) bool {
	res := db.Create(user)

	return res.Error != nil
}
