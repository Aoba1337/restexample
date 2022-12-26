package main

import (
	"github.com/aoba1337/restexample/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/users", func(ctx *gin.Context) {
		controller, err := controllers.CreateUserController()
		if err != nil {
			ctx.AbortWithStatus(500)
		} else {
			controller.GetUsers(ctx)
		}
	})

	engine.Run("localhost:11337")
}
