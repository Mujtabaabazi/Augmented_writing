package models

import (
	"github.com/Augmented_writing/wstud-augmented-writing-ss19/augmented.writing/webapp/persistence"
	"github.com/jinzhu/gorm"
)
func init()  {

}
type UserDao struct {
	db *gorm.DB
}
type Users struct {
	Id        int64  `json:"id" orm:"pk"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}

func (u * UserDao) RegisterUser(user *Users) error{
	err := u.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func NewUserDao() *UserDao {
	return &UserDao{
		db: persistence.Connection(),
	}
}