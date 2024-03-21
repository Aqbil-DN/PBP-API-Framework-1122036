package main

import (
	"github.com/gin-gonic/gin"

	"GIN/controller"
)

func main() {
	router := gin.Default()
	trustedProxies := []string{"192.168.1.2"}
	router.SetTrustedProxies(trustedProxies)
	router.GET("/GetAllUsers", controller.GinGetAllUsers)
	router.POST("/InsertUser", controller.GinInsertNewUser)
	router.PUT("/UpdateUser/:id", controller.GinUpdateUser)
	router.DELETE("/DeleteUser/:id", controller.GinDeleteUser)

	router.Run(":8888")
}
