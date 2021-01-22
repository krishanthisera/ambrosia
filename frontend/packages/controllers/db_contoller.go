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

func FindTopic(c *gin.Context) {
	var topic []models.Topic
	if err := models.DB.Table("stories").Where("name = ?", c.Param("name")).First(&topic).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": topic})
}
