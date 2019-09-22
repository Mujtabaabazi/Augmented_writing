package handlers

import (
	"github.com/Augmented_writing/wstud-augmented-writing-ss19/augmented.writing/webapp/business"
	"github.com/Augmented_writing/wstud-augmented-writing-ss19/augmented.writing/webapp/helpers"
	"github.com/Augmented_writing/wstud-augmented-writing-ss19/augmented.writing/webapp/models"
	"github.com/Augmented_writing/wstud-augmented-writing-ss19/augmented.writing/webapp/persistence"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
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

func Session(c *gin.Context) {

	sessionId, _ := c.Cookie("augmented-writing-session")
	user := &models.Users{}
	if sessionId != "" {
		session := sessions.Default(c)
		email := session.Get("email")
		if email == nil {
			c.JSON(http.StatusBadRequest, &Response{Message: "Invalid session", Data:nil})
			return
		}
		user.Email = email.(string)
		if err := persistence.Connection().Where(user).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, &Response{Message: "Invalid session", Data:nil})
			return
		}
	}
	c.JSON(http.StatusBadRequest, &Response{Message: "Ok", Data:user})
}

func ProcessRegistartion(c *gin.Context)  {

	firstname := c.PostForm("first_name")
	lastname := c.PostForm("last_name")
	username := c.PostForm("user_name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	confrimpassword := c.PostForm("password_confirm")

	_firstname, _lastname, _username, _email, _password, _confrimpassword := false, false, false, false, false, false
	_firstname = !helpers.IsEmpty(firstname)
	_lastname = !helpers.IsEmpty(lastname)
	_username = !helpers.IsEmpty(username)
	_email = !helpers.IsEmpty(email)
	_password = !helpers.IsEmpty(password)
	_confrimpassword = !helpers.IsEmpty(confrimpassword) && password == confrimpassword

	if _firstname && _lastname && _username && _email && _password && _confrimpassword {
		err := business.Registration(username, email, firstname, lastname, password)
		if err != nil {
			c.HTML(
				http.StatusBadRequest,
				"index.html", // page name
				gin.H{
					"title": "Login Page", // you can pass any number of key, values to the html page
				},
			)
			//c.JSON(http.StatusBadRequest, &Response{Message: "failed to register user"})
			return
		}
	} else {
		c.HTML(
			http.StatusBadRequest,
			"index.html", // page name
			gin.H{
				"title": "Registration Page", // you can pass any number of key, values to the html page
			},
		)
		//c.JSON(http.StatusBadRequest, &Response{Message: "Invalid registration request"})
		return
	}

	c.HTML(
		http.StatusOK,
		"home.html", // page name
		gin.H{
			"title": "Home Page", // you can pass any number of key, values to the html page
		},
	)
	//c.JSON(http.StatusOK, &Response{Message: "Register Successfully"})

}

func Logout(c *gin.Context)  {
	sessionId, _ := c.Cookie("augmented-writing-session")
	if sessionId != "" {
		session := sessions.Default(c)

		if email := session.Get("email"); email == nil {
			c.HTML(
				http.StatusBadRequest,
				"home.html", // page name
				gin.H{
					"title": "Home Page", // you can pass any number of key, values to the html page
				},
			)
			return
		}

		session.Clear()
		if err := session.Save(); err != nil {
			c.HTML(
				http.StatusBadRequest,
				"home.html", // page name
				gin.H{
					"title": "Home Page", // you can pass any number of key, values to the html page
				},
			)
			return
		}
		c.HTML(
			http.StatusOK,
			"index.html", // page name
			gin.H{
				"title": "Login Page", // you can pass any number of key, values to the html page
			},
		)
	}
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
			c.HTML(
				http.StatusUnauthorized,
				"index.html", // page name
				gin.H{
					"title": "Login Page", // you can pass any number of key, values to the html page
				},
			)
			//c.JSON(http.StatusBadRequest, &Response{Message: "Invalid username password"})
			return
		}
		// Create a session
		s := sessions.Default(c)
		s.Set("email", user.Email)

		// Save the session
		if err = s.Save(); err != nil {
			c.HTML(
				http.StatusUnauthorized,
				"index.html", // page name
				gin.H{
					"title": "Login Page", // you can pass any number of key, values to the html page
				},
			)
			//c.JSON(http.StatusBadRequest, &Response{Message: "Failed to save session"})
			return
		}

	} else {
		c.HTML(
			http.StatusBadRequest,
			"index.html", // page name
			gin.H{
				"title": "Login Page", // you can pass any number of key, values to the html page
			},
		)
		//c.JSON(http.StatusBadRequest, &Response{Message: "Username / Password is required"})
		return
	}
	if (err == nil){
		c.HTML(
			http.StatusOK,
			"home.html", // page name
			gin.H{
				"title": "Home Page", // you can pass any number of key, values to the html page
			},
		)
		//c.JSON(http.StatusOK, &Response{Message: fmt.Sprintf("Welcome %s %s", user.FirstName, user.LastName)})
	}

}
