package domain

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}

type Account struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Type   string `json:"type"`
}

type UserRepository interface {
	FindAll() ([]User, error)
	Insert(user User) error
}

type UserService interface {
	GetAllUsers() ([]User, error)
	Register(user User) error
}
