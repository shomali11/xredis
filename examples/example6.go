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
	fmt.Println(client.Expire("name", 1))
	fmt.Println(client.Expire("unknown", 1))
	fmt.Println(client.Keys("*"))
	fmt.Println(client.Get("name"))
	fmt.Println(client.Exists("name"))
	fmt.Println(client.Del("name"))
	fmt.Println(client.Exists("name"))
	fmt.Println(client.Get("name"))
	fmt.Println(client.Del("name"))
	fmt.Println(client.Append("name", "a"))
	fmt.Println(client.Append("name", "b"))
	fmt.Println(client.Append("name", "c"))
	fmt.Println(client.Get("name"))
	fmt.Println(client.GetRange("name", 0, 1))
	fmt.Println(client.Del("name"))
}
