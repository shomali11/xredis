package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.HSet("hash", "integer", "10"))
	fmt.Println(client.HSet("hash", "float", "5.5"))

	fmt.Println(client.HIncr("hash", "integer"))
	fmt.Println(client.HIncrBy("hash", "integer", 10))
	fmt.Println(client.HDecrBy("hash", "integer", 5))
	fmt.Println(client.HDecr("hash", "integer"))

	fmt.Println(client.HIncrByFloat("hash", "float", 3.3))
	fmt.Println(client.HDecrByFloat("hash", "float", 1.1))

	fmt.Println(client.HDel("hash", "integer", "float"))
}
