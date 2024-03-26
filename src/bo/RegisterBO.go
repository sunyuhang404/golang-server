package bo

type RegisterBO struct {
	// 用户名
	Username string `json:"username" binding:"required"`

	// 密码
	Password string `json:"password" binding:"required"`

	// 手机号
	Phone string `json:"phone"`

	// 头像
	Icon string `json:"icon"`
}
