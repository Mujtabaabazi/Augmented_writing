package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/Augmented_writing/wstud-augmented-writing-ss19/augmented.writing/webapp/handlers"
	"github.com/Augmented_writing/wstud-augmented-writing-ss19/augmented.writing/webapp/sessionmanagement"
)

/*
initializeRoutes (do not change)
*/
func Init(router *gin.Engine) {

	fmt.Println("Called init for routes")
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	sessionmanagement.Init(router)
	loadStaticFolders(router, false)
	initializeRoutes(router)

	router.Run()
}

/**
Add all the routes here
**/
func initializeRoutes(router *gin.Engine) {

	router.GET("/", handlers.ShowLogin)
	router.GET("/registration", handlers.ShowRegistration)
	router.POST("/processRegisteration", handlers.ProcessRegistartion)
	router.POST("/login", handlers.Login)
	router.POST("/logout", handlers.Logout)
	router.GET("/session", handlers.Session)

}

/**
ignore
**/


func loadStaticFolders(router *gin.Engine, isModule bool) {
	staticLinks := []struct {
		link string
		path string
	}{{"/css", "css"}, {"/images", "images"}, {"/js", "js"}, {"/config", "config"}}
	for i, v := range staticLinks {
		if isModule {

			staticLinks[i].path = "./../../" + v.path
		}
		router.Static(staticLinks[i].link, staticLinks[i].path)
	}

}
