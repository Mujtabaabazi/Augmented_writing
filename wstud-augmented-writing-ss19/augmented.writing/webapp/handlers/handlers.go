package handlers

import (
	"fmt"
	"github.com/Augmented_writing/wstud-augmented-writing-ss19/augmented.writing/webapp/business"
	"github.com/Augmented_writing/wstud-augmented-writing-ss19/augmented.writing/webapp/helpers"
	"github.com/Augmented_writing/wstud-augmented-writing-ss19/augmented.writing/webapp/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
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
		err := business.Registration(username, email, firstname, lastname, password)
		if err != nil {
			c.JSON(http.StatusBadRequest, &Response{Message: "failed to register user"})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, &Response{Message: "Invalid registration request"})
		return	}
	c.JSON(http.StatusOK, &Response{Message: "Register Successfully"})

}

func Login(c *gin.Context)  {

	username := c.PostForm("user_name")
	password := c.PostForm("password")

	_username, _password:= false, false
	_username = !helpers.IsEmpty(username)
	_password = !helpers.IsEmpty(password)

	var user models.Users
	var err error
	if  _username && _password {
		user, err = business.Login(username, password)
		if err != nil {
			c.JSON(http.StatusBadRequest, &Response{Message: "Invalid username password"})
			return
		}
		// Create a session
		s := sessions.Default(c)
		s.Set("email", user.Email)

		// Save the session
		if err = s.Save(); err != nil {
			c.JSON(http.StatusBadRequest, &Response{Message: "Failed to save session"})
			return
		}

	} else {
		c.JSON(http.StatusBadRequest, &Response{Message: "Username / Password is required"})
		return
	}
	if (err == nil){
	c.JSON(http.StatusOK, &Response{Message: fmt.Sprintf("Welcome %s %s", user.FirstName, user.LastName)})
	}

}
