package controllers

import (
	"os"
	"time"

	"github.com/krishanthisera/ambrosia.git/packages/models"
)

func CreateTopic(input *models.Topic) {
	// Create book
	t := time.Now()
	topic := models.Topic{Name: input.Name,
		Description: input.Description,
		Creation:    t.Format("2006:01:02 15:04:05"),
		LastEdit:    t.Format("2006:01:02 15:04:05"),
		Stories:     input.Stories,
	}
	models.DB.Table(os.Getenv("AMB_DB_TBL")).Create(&topic)
}
