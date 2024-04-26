package domain

type User struct {
	UserId       string `json:"user_id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	Email        string `json:"email"`
	FullName     string `json:"fullname"`
	PhoneNumber  string `json:"phone_number"`
	Avatar       string `json:"avatar"`
	Address      string `json:"address"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type Role struct {
	RoleID      string `json:"role_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserRole struct {
	UserId       string `json:"user_id"`
	RoleID      string `json:"role_id"`
}
