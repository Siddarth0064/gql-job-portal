package repository

import (
	"errors"
	custommodel "project-gql/models"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

// =============== NEW REPO IS USED TO INITIALIZE TO THE REPO STRUCT ============================
func NewRepo(db *gorm.DB) (*Repo, error) {
	if db == nil {
		return nil, errors.New("db connection not given")
	}

	return &Repo{db: db}, nil

}

// =============== USERS INTERFACE ==========================================
type Users interface {
	CreateUser(custommodel.User) (custommodel.User, error)
	FetchUserByEmail(string) (*custommodel.User, error)
}

// ======================== CREATE USER FUNC IS USED TO CREATE USER INFORMATION IN THE DATABASE ===================
func (r *Repo) CreateUser(u custommodel.User) (custommodel.User, error) {

	err := r.db.Create(&u).Error
	if err != nil {
		return custommodel.User{}, err
	}
	return u, nil
}

// =============== FETCH USER BY EMAIL FUNC IS USED TO RETRIVE THE USER INFORMATION IN THE DATABASE ================
func (r *Repo) FetchUserByEmail(s string) (*custommodel.User, error) {
	var u custommodel.User
	tx := r.db.Where("email=?", s).First(&u)
	if tx.Error != nil {
		return nil, nil
	}
	return &u, nil

}
