package models

type UserModel struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
}
