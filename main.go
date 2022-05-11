package main

import (
	"gin-berry/controllers"
	"gin-berry/core"
	"gin-berry/db"
	"gin-berry/models"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// set up the database
	conn := db.Init()
	models.Migrate(conn)

	// setup service
	service := core.New(func(context *gin.Context) {
		log.Println("Initial service middleware")
		context.Next()
	})
	service.Use(func(context *gin.Context) {
		log.Println("First Service middleware")
		context.Next()
	})
	service.Route("GET", "/", controllers.ServiceIndex())

	user := service.Group("/user")
	user.Use(func(context *gin.Context) {
		log.Println("Generic User service middleware")
		context.Next()
	})
	//
	//user.Use(func(context *gin.Context) {
	//	log.Println("User Service Middleware 2", context.MustGet("routeConfig"))
	//	context.Next()
	//})
	//user.Use(func(context *gin.Context) {
	//	log.Println("User Service Middleware 3", context.MustGet("routeConfig"))
	//	context.Next()
	//})
	//
	user.Route("GET", "/", controllers.ServiceIndex())
	user.Route("GET", "/detail", controllers.ServiceIndex())
	//user.Route("GET", "/first", user.GetUserDetail())

	err := service.Run(":5002")
	if err != nil {
		log.Fatal("Failed to start server. \n", err)
	}
}
