# xredis [![Go Report Card](https://goreportcard.com/badge/github.com/shomali11/xredis)](https://goreportcard.com/report/github.com/shomali11/xredis) [![GoDoc](https://godoc.org/github.com/shomali11/xredis?status.svg)](https://godoc.org/github.com/shomali11/xredis) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Built on top of [github.com/garyburd/redigo](https://github.com/garyburd/redigo) with the idea to simplify creating a Redis client, provide type safe calls and encapsulate the low level details to easily integrate with Redis.

## Features

* Type safe client
* Easy to setup using
    * Default client
    * Custom client via set options
    * `redigo`'s `redis.Pool`
* Connection pool provided automatically
* Support for Redis Sentinel
    * Writes go to the Master
    * Reads go to the Slaves. Falls back on Master if none are available.
* Supports the following Redis commands
    * **ECHO**, **INFO**, **PING**, **FLUSH**, **FLUSHALL**, **EXPIRE**, **APPEND**
    * **SET**, **SETEX**, **SETNX**, **GET**, **DEL**, **EXISTS**, **KEYS**, **SCAN**, **GETRANGE**, **SETRANGE**
    * **HSET**, **HGET**, **HGETALL**, **HDEL**, **HEXISTS**, **HKEYS**, **HSCAN**
    * **INCR**, **INCRBY**, **INCRBYFLOAT**, **DECR**, **DECRBY**, **DECRBYFLOAT**
    * **HINCR**, **HINCRBY**, **HINCRBYFLOAT**, **HDECR**, **HDECRBY**, **HDECRBYFLOAT**
    * _More coming soon_
* Full access to Redigo's API [github.com/garyburd/redigo](https://github.com/garyburd/redigo)

## Usage

Using `govendor` [github.com/kardianos/govendor](https://github.com/kardianos/govendor):

```
govendor fetch github.com/shomali11/xredis
```

## Dependencies

* `redigo` [github.com/garyburd/redigo](https://github.com/garyburd/redigo)
* `go-sentinel` [github.com/FZambia/go-sentinel](https://github.com/FZambia/go-sentinel)

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
defaultConnectionMaxActive   = 10000
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

Using `SetupSentinelClient` to create a redis sentinel client using provided options

```go
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

	fmt.Println(client.Ping()) // PONG <nil>
}
```

Available options to set

```go
type SentinelOptions struct {
	Addresses             []string
	MasterName            string
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
}
```

## Example 4

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

## Example 5

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
	fmt.Println(client.FlushDb())      // <nil>
	fmt.Println(client.FlushAll())     // <nil>
	fmt.Println(client.Info())         
}
```

## Example 6

Using the `Set`, `Keys`, `Get`, `Exists`, `Expire`, `Append`, `GetRange`, `SetRange` and `Del` commands to show how to set, get and delete keys and values.
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

	fmt.Println(client.Set("name", "Raed Shomali")) // true <nil>
	fmt.Println(client.SetNx("name", "Hello"))      // false <nil>
	fmt.Println(client.SetEx("id", "10", 1))        // true <nil>
	fmt.Println(client.Expire("name", 1))           // true <nil>
	fmt.Println(client.Expire("unknown", 1))        // false <nil>
	fmt.Println(client.Keys("*"))                   // [id name] <nil>
	fmt.Println(client.Get("name"))                 // "Raed Shomali" true <nil>
	fmt.Println(client.Exists("name"))              // true <nil>
	fmt.Println(client.Del("name"))                 // 1 <nil>
	fmt.Println(client.Exists("name"))              // false <nil>
	fmt.Println(client.Get("name"))                 // "" false <nil>
	fmt.Println(client.Del("name"))                 // 0 <nil>
	fmt.Println(client.Append("name", "a"))         // 1 <nil>
	fmt.Println(client.Append("name", "b"))         // 2 <nil>
	fmt.Println(client.Append("name", "c"))         // 3 <nil>
	fmt.Println(client.Get("name"))                 // "abc" true <nil>
	fmt.Println(client.GetRange("name", 0 , 1))     // "ab" <nil>
	fmt.Println(client.SetRange("name", 2, "xyz"))  // 5 <nil>
	fmt.Println(client.Get("name"))                 // "abxyz" <nil>
	fmt.Println(client.Scan(0, "*"))                // 0 [name id] <nil>
	fmt.Println(client.Del("id", "name"))           // 2 <nil>
}
```

## Example 7

Using the `Incr`, `IncrBy`, `IncrByFloat`, `Decr`, `DecrBy`, `DecrByFloat` commands, we can increment and decrement a key's value

```go
package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.Set("integer", "10"))       // true <nil>
	fmt.Println(client.Set("float", "5.5"))        // true <nil>

	fmt.Println(client.Get("integer"))             // 10 true <nil>
	fmt.Println(client.Get("float"))               // 5.5 true <nil>

	fmt.Println(client.Incr("integer"))            // 11 <nil>
	fmt.Println(client.IncrBy("integer", 10))      // 21 <nil>
	fmt.Println(client.DecrBy("integer", 5))       // 16 <nil>
	fmt.Println(client.Decr("integer"))            // 15 <nil>

	fmt.Println(client.IncrByFloat("float", 3.3))  // 8.8 <nil>
	fmt.Println(client.DecrByFloat("float", 1.1))  // 7.7 <nil>

	fmt.Println(client.Get("integer"))             // 15 true <nil>
	fmt.Println(client.Get("float"))               // 7.7 true <nil>

	fmt.Println(client.Del("integer", "float"))    // 2 <nil>
}
```

## Example 8

Using the `HSet`, `HKeys`, `HGet`, `HGetAll`, `HExists` and `HDel` commands to show how to set, get and delete hash keys, fields and values.
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

	fmt.Println(client.HSet("hash", "name", "Raed Shomali")) // true <nil>
	fmt.Println(client.HSet("hash", "sport", "Football"))    // true <nil>
	fmt.Println(client.HKeys("hash"))                        // [name sport] <nil>
	fmt.Println(client.HScan("hash", 0, "*"))                // 0 [name Raed Shomali sport Football] <nil>
	fmt.Println(client.HGet("hash", "name"))                 // "Raed Shomali" true <nil>
	fmt.Println(client.HGetAll("hash"))                      // map[name:Raed Shomali sport:Football] <nil>
	fmt.Println(client.HExists("hash", "name"))              // true <nil>
	fmt.Println(client.HDel("hash", "name", "sport"))        // 2 <nil>
	fmt.Println(client.HGet("hash", "name"))                 // "" false <nil>
	fmt.Println(client.HExists("hash", "name"))              // false <nil>
	fmt.Println(client.HGetAll("hash"))                      // map[] nil
	fmt.Println(client.HDel("hash", "name"))                 // 0 <nil>
	fmt.Println(client.HKeys("hash"))                        // [] <nil>
}
```

## Example 9

Using the `HIncr`, `HIncrBy`, `HIncrByFloat`,`HDecr`, `HDecrBy` and `HDecrByFloat` commands to show how to increment and decrement hash fields' values.

```go
package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.HSet("hash", "integer", "10"))       // true <nil>
	fmt.Println(client.HSet("hash", "float", "5.5"))        // true <nil>

	fmt.Println(client.HIncr("hash", "integer"))            // 11 <nil>
	fmt.Println(client.HIncrBy("hash", "integer", 10))      // 21 <nil>
	fmt.Println(client.HDecrBy("hash", "integer", 5))       // 16 <nil>
	fmt.Println(client.HDecr("hash", "integer"))            // 15 <nil>

	fmt.Println(client.HIncrByFloat("hash", "float", 3.3))  // 8.8 <nil>
	fmt.Println(client.HDecrByFloat("hash", "float", 1.1))  // 7.7 <nil>

	fmt.Println(client.HDel("hash", "integer", "float"))    // 2 <nil>
}
```

## Example 10

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