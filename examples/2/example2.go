package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	options := &xredis.Options{
		Host: "localhost",
		Port: 6379,
	}

	client := xredis.SetupClient(options)
	defer client.Close()

	fmt.Println(client.Ping())
}
