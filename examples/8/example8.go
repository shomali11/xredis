package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.HSet("hash", "name", "Raed Shomali"))
	fmt.Println(client.HSet("hash", "sport", "Football"))
	fmt.Println(client.HKeys("hash"))
	fmt.Println(client.HScan("hash", 0, "*"))
	fmt.Println(client.HGet("hash", "name"))
	fmt.Println(client.HGetAll("hash"))
	fmt.Println(client.HExists("hash", "name"))
	fmt.Println(client.HDel("hash", "name", "sport"))
	fmt.Println(client.HGet("hash", "name"))
	fmt.Println(client.HExists("hash", "name"))
	fmt.Println(client.HGetAll("hash"))
	fmt.Println(client.HDel("hash", "name"))
	fmt.Println(client.HKeys("hash"))
}
