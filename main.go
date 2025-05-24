package main

import (
	"github.com/Divyshekhar/go-oauth/controllers"
	"github.com/Divyshekhar/go-oauth/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitOAuth()
	r := gin.Default()
	r.GET("/login", controllers.LoginHandler)
	r.GET("/oauth2callback", controllers.CallBackHandler)
	r.GET("/emails", controllers.EmailHandler)

	r.Run(":8080")

}
