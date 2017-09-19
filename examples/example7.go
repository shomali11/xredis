package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.Set("integer", "10"))
	fmt.Println(client.Set("float", "5.5"))

	fmt.Println(client.Get("integer"))
	fmt.Println(client.Get("float"))

	fmt.Println(client.Incr("integer"))
	fmt.Println(client.IncrBy("integer", 10))
	fmt.Println(client.DecrBy("integer", 5))
	fmt.Println(client.Decr("integer"))

	fmt.Println(client.IncrByFloat("float", 3.3))
	fmt.Println(client.DecrByFloat("float", 1.1))

	fmt.Println(client.Get("integer"))
	fmt.Println(client.Get("float"))

	fmt.Println(client.Del("integer", "float"))
}
