package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishanthisera/ambrosia.git/packages/models"
)

func FindTopics(c *gin.Context) {
	var topic []models.Topic
	models.DB.Table("stories").Find(&topic)
	fmt.Println(topic)
	c.JSON(http.StatusOK, gin.H{"data": topic})
}
