# xredis [![Go Report Card](https://goreportcard.com/badge/github.com/shomali11/xredis)](https://goreportcard.com/report/github.com/shomali11/xredis) [![GoDoc](https://godoc.org/github.com/shomali11/xredis?status.svg)](https://godoc.org/github.com/shomali11/xredis) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`xredis` is a wrapper around [redigo](https://github.com/garyburd/redigo)

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

	fmt.Println(client.Ping())
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

	fmt.Println(client.Ping())
}
```

Available options to modify

```go
// Options contains redis options
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

	fmt.Println(client.Ping())
}

```