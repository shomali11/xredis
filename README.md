# xredis [![Go Report Card](https://goreportcard.com/badge/github.com/shomali11/xredis)](https://goreportcard.com/report/github.com/shomali11/xredis) [![GoDoc](https://godoc.org/github.com/shomali11/xredis?status.svg)](https://godoc.org/github.com/shomali11/xredis) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Built on top of [github.com/garyburd/redigo](https://github.com/garyburd/redigo) with the idea to simplify creating a Redis client, provide type safe calls and encapsulate the low level details to easily integrate with Redis.

## Features

* Type safe client
* Easy to setup using
    * Default client
    * Custom client via set options
    * `redigo`'s `redis.Pool`
* Connection pool provided automatically
* Supports the following Redis commands
    * **ECHO**, **INFO**, **PING**
    * **SET**, **GET**, **DEL**, **EXISTS**
    * **HSET**, **HGET**, **HGETALL**, **HDEL**, **HEXISTS**
    * **INCR**, **INCRBY**, **DECR**, **DECRBY**, **HINCRBY**, **HINCRBYFLOAT**,
    * _More coming soon_
* Full access to Redigo's API [github.com/garyburd/redigo](https://github.com/garyburd/redigo)

## Usage

Using `govendor` [github.com/kardianos/govendor](https://github.com/kardianos/govendor):

```
govendor fetch github.com/shomali11/xredis
```

## Dependencies

* `redigo` [github.com/garyburd/redigo](https://github.com/garyburd/redigo)

# Examples

## Example 1

Using `DefaultClient` to create a redis client with default options

```go
package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.Ping()) // PONG <nil>
}
```

List of default options

```text
defaultHost                  = "localhost"
defaultPort                  = 6379
defaultPassword              = ""
defaultDatabase              = 0
defaultNetwork               = "tcp"
defaultConnectTimeout        = time.Second
defaultWriteTimeout          = time.Second
defaultReadTimeout           = time.Second
defaultConnectionIdleTimeout = 240 * time.Second
defaultConnectionMaxIdle     = 100
defaultConnectionMaxActive   = 1000
defaultConnectionWait        = false
defaultTlsConfig             = nil
defaultTlsSkipVerify         = false
defaultTestOnBorrowTimeout   = time.Minute
```

## Example 2

Using `SetupClient` to create a redis client using provided options

```go
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

	fmt.Println(client.Ping()) // PONG <nil>
}
```

Available options to set

```go
type Options struct {
	Host                  string
	Port                  int
	Password              string
	Database              int
	Network               string
	ConnectTimeout        time.Duration
	WriteTimeout          time.Duration
	ReadTimeout           time.Duration
	ConnectionIdleTimeout time.Duration
	ConnectionMaxIdle     int
	ConnectionMaxActive   int
	ConnectionWait        bool
	TlsConfig             *tls.Config
	TlsSkipVerify         bool
	TestOnBorrowPeriod    time.Duration
}
```

## Example 3

Using `NewClient` to create a redis client using `redigo`'s `redis.Pool`

```go
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/shomali11/xredis"
)

func main() {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	client := xredis.NewClient(pool)
	defer client.Close()

	fmt.Println(client.Ping()) // PONG <nil>
}
```

## Example 4

Using the `Ping`, `Echo` & `Info` commands to ping, echo messages and return redis' information and statistics

```go
package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.Ping())         // PONG <nil>
	fmt.Println(client.Echo("Hello"))  // Hello <nil>
	fmt.Println(client.Info())         
}
```

## Example 5

Using the `Set`, `Get`, `Exists` and `Del` commands to show how to set, get and delete keys and values.
_Note that the `Get` returns 3 values, a `string` result, a `bool` that determines whether the key exists and an `error`_

```go
package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.Set("name", "Raed Shomali")) // OK <nil>
	fmt.Println(client.Get("name"))                 // "Raed Shomali" true <nil>
	fmt.Println(client.Exists("name"))              // true <nil>
	fmt.Println(client.Del("name"))                 // 1 <nil>
	fmt.Println(client.Exists("name"))              // false <nil>
	fmt.Println(client.Get("name"))                 // "" false <nil>
	fmt.Println(client.Del("name"))                 // 0 <nil>
}
```

## Example 6

Using the `Incr` command, we can increment a key's value by one

```go
package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.Set("number", "10")) // OK <nil>
	fmt.Println(client.Get("number"))       // 10 true <nil>
	fmt.Println(client.Incr("number"))      // 11 <nil>
	fmt.Println(client.Get("number"))       // 11 true <nil>
	fmt.Println(client.Del("number"))       // 1 <nil>
}
```

## Example 7

Using the `IncrBy` command, we can increment a key's value by an increment

```go
package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.Set("number", "10"))  // OK <nil>
	fmt.Println(client.Get("number"))        // 10 true <nil>
	fmt.Println(client.IncrBy("number", 10)) // 20 <nil> 
	fmt.Println(client.Get("number"))        // 20 true <nil>
	fmt.Println(client.Del("number"))        // 1 <nil>
}
```

## Example 8

Using the `Decr` command, we can decrement a key's value by one

```go
package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.Set("number", "10")) // OK <nil>
	fmt.Println(client.Get("number"))       // 10 true <nil>
	fmt.Println(client.Decr("number"))      // 9 <nil>
	fmt.Println(client.Get("number"))       // 9 true <nil>
	fmt.Println(client.Del("number"))       // 1 <nil>
}
```

## Example 9

Using the `DecrBy` command, we can decrement a key's value by an decrement

```go
package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.Set("number", "10"))  // OK <nil>
	fmt.Println(client.Get("number"))        // 10 true <nil>
	fmt.Println(client.DecrBy("number", 5))  // 5 <nil> 
	fmt.Println(client.Get("number"))        // 5 true <nil>
	fmt.Println(client.Del("number"))        // 1 <nil>
}
```

## Example 10

Using the `HSet`, `HGet`, `HGetAll`, `HExists` and `HDel` commands to show how to set, get and delete hash keys, fields and values.
_Note that the `HGetAll` returns 2 values, a `map[string]string` result and an `error`_

```go
package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.HSet("hash", "name", "Raed Shomali")) // 1 <nil>
	fmt.Println(client.HSet("hash", "sport", "Football"))    // 1 <nil>
	fmt.Println(client.HGet("hash", "name"))                 // "Raed Shomali" true <nil>
	fmt.Println(client.HGetAll("hash"))                      // map[name:Raed Shomali sport:Football] <nil>
	fmt.Println(client.HExists("hash", "name"))              // true <nil>
	fmt.Println(client.HDel("hash", "name", "sport"))        // 2 <nil>
	fmt.Println(client.HGet("hash", "name"))                 // "" false <nil>
	fmt.Println(client.HExists("hash", "name"))              // false <nil>
	fmt.Println(client.HGetAll("hash"))                      // map[] nil
	fmt.Println(client.HDel("hash", "name"))                 // 0 <nil>
}
```

## Example 11

Using the `HIncrBy` and `HIncrByFloat` commands to show how to increment hash fields by integer and float increments values.

```go
package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.HSet("hash", "integer", "10"))      // 1 <nil>
	fmt.Println(client.HSet("hash", "float", "5.5"))       // 1 <nil>
	fmt.Println(client.HIncrBy("hash", "integer", 10))     // 20 <nil>
	fmt.Println(client.HIncrByFloat("hash", "float", 3.3)) // 8.8 <nil>
	fmt.Println(client.HDel("hash", "integer", "float"))   // 2 <nil>
}
```

## Example 12

Can't find the command you want? You have full access to `redigo`'s API.

```go
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
```