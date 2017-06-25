package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.Set("name", "Raed Shomali"))
	fmt.Println(client.SetNx("name", "Hello"))
	fmt.Println(client.SetEx("id", "10", 1))
	fmt.Println(client.Keys("*"))
	fmt.Println(client.Get("name"))
	fmt.Println(client.Exists("name"))
	fmt.Println(client.Del("name"))
	fmt.Println(client.Exists("name"))
	fmt.Println(client.Get("name"))
	fmt.Println(client.Del("name"))
}
