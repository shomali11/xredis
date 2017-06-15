package xredis

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	pingCommand = "PING"
)

// DefaultClient returns a client with default options
func DefaultClient() *Client {
	pool := newPool(&Options{})
	return &Client{pool: pool}
}

// SetupClient returns a client with provided options
func SetupClient(options *Options) *Client {
	pool := newPool(options)
	return &Client{pool: pool}
}

// NewClient returns a client using provided redis.Pool
func NewClient(pool *redis.Pool) *Client {
	return &Client{pool: pool}
}

// Client redis client
type Client struct {
	pool *redis.Pool
}

// GetConnection gets a connection
func (r *Client) GetConnection() redis.Conn {
	return r.pool.Get()
}

func newPool(options *Options) *redis.Pool {
	connectionIdleTimeout := options.GetConnectionIdleTimeout()
	connectionMaxActive := options.GetConnectionMaxActive()
	connectionMaxIdle := options.GetConnectionMaxIdle()
	connectionWait := options.GetConnectionWait()

	return &redis.Pool{
		IdleTimeout:  connectionIdleTimeout,
		MaxActive:    connectionMaxActive,
		MaxIdle:      connectionMaxIdle,
		Wait:         connectionWait,
		Dial:         dial(options),
		TestOnBorrow: testOnBorrow(options),
	}
}

func dial(options *Options) func() (redis.Conn, error) {
	network := options.GetNetwork()
	address := options.GetAddress()

	dialOptions := make([]redis.DialOption, 7)
	dialOptions[0] = redis.DialPassword(options.GetPassword())
	dialOptions[1] = redis.DialDatabase(options.GetDatabase())
	dialOptions[2] = redis.DialConnectTimeout(options.GetConnectTimeout())
	dialOptions[3] = redis.DialWriteTimeout(options.GetWriteTimeout())
	dialOptions[4] = redis.DialReadTimeout(options.GetReadTimeout())
	dialOptions[5] = redis.DialTLSSkipVerify(options.GetTlsSkipVerify())
	dialOptions[6] = redis.DialTLSConfig(options.GetTlsConfig())

	return func() (redis.Conn, error) {
		connection, err := redis.Dial(network, address, dialOptions...)
		if err != nil {
			return nil, err
		}
		return connection, nil
	}
}

func testOnBorrow(options *Options) func(redis.Conn, time.Time) error {
	period := options.GetTestOnBorrowPeriod()

	return func(connection redis.Conn, t time.Time) error {
		if time.Since(t) < period {
			return nil
		}

		_, err := connection.Do(pingCommand)
		return err
	}
}
