package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"wstud-augmented-writing-ss19/augmented.writing/webapp/routes"
)

var router *gin.Engine

func main() {

	routes.Init(router)
	log.Print(" routes initialized...")

}
