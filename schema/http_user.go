package schema

type RegisterRequest struct {
	Username string `json:"username" binding:"required" example:"test_user"`
	Password string `json:"password" binding:"required" example:"123456"`
}

// 用户信息
type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}
