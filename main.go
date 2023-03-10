package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rudiath95/blog-rest-API/controllers"
	"github.com/rudiath95/blog-rest-API/ini"
	"github.com/rudiath95/blog-rest-API/middleware"
)

func init() {
	ini.LoadEnvVariables()
	ini.ConnecttoDB()
	ini.SyncDatabases()
}

func main() {
	r := gin.Default()
	//>>Start Route

	//#User
	r.GET("/getuser", middleware.RequiredAuth, controllers.GetUser)
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.PUT("/edit-user/:id", middleware.RequiredAuth, controllers.EditUser)
	r.DELETE("/delete-user/:id", middleware.RequiredAuth, controllers.DeleteUser)

	//#Blog

	//>>End Route
	r.Run()
}
