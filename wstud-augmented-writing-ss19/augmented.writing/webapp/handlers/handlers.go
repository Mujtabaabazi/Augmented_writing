package handlers

import (
	"fmt"
	"github.com/Augmented_writing/wstud-augmented-writing-ss19/augmented.writing/webapp/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"

	"github.com/Augmented_writing/wstud-augmented-writing-ss19/augmented.writing/webapp/helpers"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status int64
	Message string
	Data interface{}
}

func ShowLogin(c *gin.Context) {

	// Here is a sample on how to redirect to a page.
	c.HTML(
		http.StatusOK,
		"index.html", // page name
		gin.H{
			"title": "", // you can pass any number of key, values to the huml page
		},
	)
}

func ShowRegistration(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"index.html", // page name
		gin.H{
			"title": "Registration Page", // you can pass any number of key, values to the html page
		},
	)

}

func ProcessRegistartion(c *gin.Context)  {

	firstname := c.PostForm("firstnamesignup")
	lastname := c.PostForm("lastnamesignup")
	username := c.PostForm("usernamesignup")
	email := c.PostForm("emailsignup")
	password := c.PostForm("passwordsignup")
	confrimpassword := c.PostForm("passwordsignup_confirm")

	_firstname, _lastname, _username, _email, _password, _confrimpassword := false, false, false, false, false, false
	_firstname = !helpers.IsEmpty(firstname)
	_lastname = !helpers.IsEmpty(lastname)
	_username = !helpers.IsEmpty(username)
	_email = !helpers.IsEmpty(email)
	_password = !helpers.IsEmpty(password)
	_confrimpassword = !helpers.IsEmpty(confrimpassword)

	if _firstname && _lastname && _username && _email && _password && _confrimpassword {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			//return err
		}
		userDao := models.NewUserDao()
		err = userDao.RegisterUser(&models.Users{
			FirstName: firstname,
			LastName: lastname,
			Email: email,
			UserName: username,
			Password: string(hashedPassword),
		})
		if err != nil {
			//return err
		}
		//return nil
	} else {
		fmt.Println(c, "This fields can not be blank!")
	}
	c.JSON(http.StatusOK, &Response{Message: "UserUpdationSuccessful"})

}
