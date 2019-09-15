package sessionmanagement

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const sessionName = "augmented-writing-session"

/*
Initialize Sessions
*/
func Init(router *gin.Engine) {
	//s := sessions.NewCookieStore([]byte("secret"))
	store := cookie.NewStore([]byte("secret"))

	router.Use(sessions.Sessions(sessionName, store))

}
