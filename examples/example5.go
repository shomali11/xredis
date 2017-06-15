package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.Set("name", "Raed Shomali"))
	fmt.Println(client.Get("name"))
	fmt.Println(client.Del("name"))
	fmt.Println(client.Get("name"))
	fmt.Println(client.Del("name"))
}
