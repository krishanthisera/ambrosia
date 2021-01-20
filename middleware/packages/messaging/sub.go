package messaging

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/krishanthisera/ambrosia.git/packages/models"
)

func RedSub(channel, conString string) {
	conn, conErr := redis.Dial("tcp", "localhost:6379")
	if conErr != nil {
		panic("Redis Dial")
	}
	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe("topics")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			var topic models.Topic
			err := json.Unmarshal(v.Data, &topic)
			if err != nil {
				panic("Unmarshal")
			}
			DbPush(topic)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
		}
	}
	defer conn.Close()
}

func DbPush(topic models.Topic) {
	dbString := new(models.PsqlConnect)
	dbString.Dbname = "topics"
	dbString.Port = 5432
	dbString.Password = "123"
	dbString.User = "user"
	dbString.Host = "localhost"

	models.DbInsert(models.DbConnect(dbString), topic)

}
