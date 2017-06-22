package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	connection := client.GetConnection()
	defer connection.Close()

	fmt.Println(redis.String(connection.Do("INFO")))
}
