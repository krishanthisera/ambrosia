package messaging

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/krishanthisera/ambrosia.git/packages/controllers"
	"github.com/krishanthisera/ambrosia.git/packages/models"
)

type TopicInput struct {
	Id          int    `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Creation    string `json:"creation"`
	LastEdit    string `json:"lastedit"`
	Stories     string `json:"stories"`
}

func RedSub(channel, conString string) {
	conn, conErr := redis.Dial("tcp", conString)
	if conErr != nil {
		panic("Redis Dial")
	}
	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe(channel)
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			var topic models.Topic
			err := json.Unmarshal(v.Data, &topic)
			if err != nil {
				panic("Unmarshal")
			}
			controllers.CreateTopic(&topic)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
		}
	}
	defer conn.Close()
}
