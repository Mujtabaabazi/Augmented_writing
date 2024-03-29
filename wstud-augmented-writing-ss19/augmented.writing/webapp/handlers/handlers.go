package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"wstud-augmented-writing-ss19/augmented.writing/webapp/helpers"
)

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

func ProcessRegistartion(c *gin.Context) {

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
		fmt.Fprintln(c, "Firstname for Register : ", firstname)
		fmt.Fprintln(c, "Lastname for Register : ", lastname)
		fmt.Fprintln(c, "Username for Register : ", username)
		fmt.Fprintln(c, "Email for Register : ", email)
		fmt.Fprintln(c, "Password for Register : ", password)
		fmt.Fprintln(c, "ConfirmPassword for Register : ", confrimpassword)

	} else {
		fmt.Fprintln(w, "This fields can not be blank!")
	}

}
