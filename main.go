package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	docs "maily/go-backend/docs"
	"maily/go-backend/src/database"
	"maily/go-backend/src/routers"
)

var host string = "0.0.0.0"

// var host string = "localhost"
var port string = "8090"

// @title        Maily API
// @version      1.0
// @description  Maily backend server

// @host      localhost:8090
// @BasePath  /api/v1
func main() {
	router := gin.Default()
	router.Use(database.Connect())

	// Swagger route
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	println("Swagger route:", router.BasePath()+"/swagger")

	routers.Init(router)
	router.Run(fmt.Sprintf("%s:%s", host, port))
}
