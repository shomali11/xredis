package eredigo

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	pingCommand = "PING"
)

func DefaultClient() *Client {
	pool := newPool(&Options{})
	return &Client{pool: pool}
}

func SetupClient(options *Options) *Client {
	pool := newPool(options)
	return &Client{pool: pool}
}

func NewClient(pool *redis.Pool) *Client {
	return &Client{pool: pool}
}

type Client struct {
	pool *redis.Pool
}

func (r *Client) GetConnection() redis.Conn {
	return r.pool.Get()
}

func newPool(options *Options) *redis.Pool {
	connectionIdleTimeout := options.ConnectionIdleTimeout()
	connectionMaxActive := options.ConnectionMaxActive()
	connectionMaxIdle := options.ConnectionMaxIdle()
	connectionWait := options.ConnectionWait()

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
	network := options.Network()
	address := options.Address()

	dialOptions := make([]redis.DialOption, 7)
	dialOptions[0] = redis.DialPassword(options.Password())
	dialOptions[1] = redis.DialDatabase(options.Database())
	dialOptions[2] = redis.DialConnectTimeout(options.ConnectTimeout())
	dialOptions[3] = redis.DialWriteTimeout(options.WriteTimeout())
	dialOptions[4] = redis.DialReadTimeout(options.ReadTimeout())
	dialOptions[5] = redis.DialTLSSkipVerify(options.TlsSkipVerify())
	dialOptions[6] = redis.DialTLSConfig(options.TlsConfig())

	return func() (redis.Conn, error) {
		connection, err := redis.Dial(network, address, dialOptions...)
		if err != nil {
			return nil, err
		}
		return connection, nil
	}
}

func testOnBorrow(options *Options) func(redis.Conn, time.Time) error {
	period := options.TestOnBorrowPeriod()

	return func(connection redis.Conn, t time.Time) error {
		if time.Since(t) < period {
			return nil
		}

		_, err := connection.Do(pingCommand)
		return err
	}
}
