package messaging

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
	"github.com/krishanthisera/ambrosia.git/packages/models"
)

func RedPub(data models.Topic, conString, channel string) {
	conn, conErr := redis.Dial("tcp", conString)
	if conErr != nil {
		panic("Redis Dial")
	}

	jsonString, jsonErr := json.Marshal(data)
	if jsonErr != nil {
		panic("Json Conversion")
	}
	_, msgErr := conn.Do("PUBLISH", channel, jsonString)
	if msgErr != nil {
		panic("Redis Subscribe")
	}
	defer conn.Close()
}
