package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/krishanthisera/ambrosia.git/packages/models"
)

func FindTopics(c *gin.Context) {
	var topic []models.Topic
	models.DB.Table(os.Getenv("AMB_DB_TBL")).Find(&topic)
	fmt.Println(topic)
	c.JSON(http.StatusOK, gin.H{"data": topic})
}

func FindTopic(c *gin.Context) {
	var topic []models.Topic
	if err := models.DB.Table(os.Getenv("AMB_DB_TBL")).Where("name = ?", c.Param("name")).First(&topic).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": topic})
}
