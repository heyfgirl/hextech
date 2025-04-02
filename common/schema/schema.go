package schema

// 通用列表返回格式
type PageResult struct {
	List     any   `json:"list"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}

// 用户信息
type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}
