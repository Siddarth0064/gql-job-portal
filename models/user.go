package custommodel

import (
	//"project/internal/database"

	"gorm.io/gorm"
)

// ======================  CREATING USER TABLE ==================================
type User struct {
	gorm.Model
	UserName     string `json:"name" validate:"required,unique" gorm:"unique,notnull"`
	Email        string `json:"email" validate:"required"`
	PasswordHash string `json:"_" validate:"required"`
}

// ================ USER SIGN IN FIELDS =========================================
type UserSignup struct {
	UserName string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// =================== USER LOGIN FIELDS ============================================
type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
