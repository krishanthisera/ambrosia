package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krishanthisera/ambrosia.git/packages/messaging"
	"github.com/krishanthisera/ambrosia.git/packages/models"
)

type TopicInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Stories     string `json:"stories" binding:"required"`
}

func CreateTopic(c *gin.Context) {

	var input TopicInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create topic
	topic := models.Topic{Name: input.Name,
		Description: input.Description,
		LastEdit:    time.Now().String(),
		Creation:    time.Now().String(),
		Stories:     input.Stories,
	}

	//Create Messaging Channel
	channel := models.RedisConnect{ConString: "localhost:6379",
		Channel: "topics-1",
	}

	// Publish the message
	messaging.RedPub(topic, channel.ConString, channel.Channel)
	c.JSON(http.StatusOK, gin.H{"data": topic})
}
