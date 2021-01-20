package main

import (
	"github.com/krishanthisera/ambrosia.git/packages/controllers"
)

func main() {
	//r := gin.Default()
	//models.ConnectDatabase()
	//r.GET("/topics", controllers.GetAllTopics) // new
	//r.POST("/topics", controllers.CreateTopic)
	//r.Run()

	controllers.PushTopic()

}
