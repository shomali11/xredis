package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	options := &xredis.SentinelOptions{
		Addresses:  []string{"localhost:26379"},
		MasterName: "master",
	}

	client := xredis.SetupSentinelClient(options)
	defer client.Close()

	fmt.Println(client.Ping())
}
