package business

import (
	"github.com/Augmented_writing/wstud-augmented-writing-ss19/augmented.writing/webapp/models"
	"github.com/Augmented_writing/wstud-augmented-writing-ss19/augmented.writing/webapp/persistence"
	"golang.org/x/crypto/bcrypt"
)

// define all SQL queries and create 1 func for each call

const (
	SelectUser = `SELECT ... FROM ... WHERE ...`
)

func Login(userName, password string) (models.Users, error) {
	u := models.Users{UserName: userName}
	if err := persistence.Connection().Where(u).First(&u).Error; err != nil {
		return u, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return u, err
	}
	return u, nil
}