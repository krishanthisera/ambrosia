package main

import (
	"github.com/gin-gonic/gin"
	"github.com/krishanthisera/ambrosia.git/packages/controllers"
	"github.com/krishanthisera/ambrosia.git/packages/models"
)

func main() {
	r := gin.Default()
	models.DbConnect()
	r.GET("/topics/:name", controllers.FindTopic)
	r.GET("/topics", controllers.FindTopics)
	r.POST("/topics", controllers.CreateTopic)
	r.Run()
}
