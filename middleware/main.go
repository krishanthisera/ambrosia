package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/krishanthisera/ambrosia.git/packages/messaging"
	"github.com/krishanthisera/ambrosia.git/packages/models"
)

func main() {
	//r := gin.Default()
	models.DbConnect()
	//r.GET("/topics", controllers.GetAllTopics) // new
	//r.POST("/topics", controllers.CreateTopic)
	//r.Run()

	//processe("localhost:6379", "topics-1", "topic-2", "topic-3")
	messaging.RedSub(os.Getenv("AMB_RED_CHAN"), os.Getenv("AMB_RED_HOST"))
}

func processe(conString string, channels ...string) {
	for _, v := range channels {
		go messaging.RedSub(v, conString)
		fmt.Println("Routine Created: ", v)
	}
	fmt.Println(getGID())

}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
