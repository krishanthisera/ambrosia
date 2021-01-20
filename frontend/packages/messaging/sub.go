package messaging

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

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
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
		}
	}
	defer conn.Close()
}
