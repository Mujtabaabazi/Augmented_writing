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

func Registration(username, email, firstname, lastname, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userDao := models.NewUserDao()
	if err = userDao.RegisterUser(&models.Users{
		FirstName: firstname,
		LastName: lastname,
		Email: email,
		UserName: username,
		Password: string(hashedPassword),
	}); err != nil {
		return err
	}


	return nil

}