package entities

type User struct {
	ID         int64  `json:"id" gorm:"primaryKey"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	IsVerified bool   `json:"is_verified" gorm:"default:false"`
	CreatedAt  int64  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  int64  `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserCredential struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserRegisteration struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserUpdation struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}
