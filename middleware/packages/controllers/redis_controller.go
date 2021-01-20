package controllers

import (
	"github.com/krishanthisera/ambrosia.git/packages/messaging"
	"github.com/krishanthisera/ambrosia.git/packages/models"
)

func PushTopic() {
	channel := models.RedisConnect{ConString: "localhost:6379",
		Channel: "topics",
	}
	messaging.RedSub(channel.ConString, channel.Channel)
}
