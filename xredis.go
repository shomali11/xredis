package xredis

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	setCommand     = "SET"
	delCommand     = "DEL"
	getCommand     = "GET"
	pingCommand    = "PING"
	echoCommand    = "ECHO"
	infoCommand    = "INFO"
	incrCommand    = "INCR"
	decrCommand    = "DECR"
	hsetCommand    = "HSET"
	hgetCommand    = "HGET"
	hdelCommand    = "HDEL"
	hgetAllCommand = "HGETALL"
	incrByCommand  = "INCRBY"
	decrByCommand  = "DECRBY"
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

// GetConnection gets a connection from the pool
func (c *Client) GetConnection() redis.Conn {
	return c.pool.Get()
}

// Ping pings redis
func (c *Client) Ping() (string, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.String(connection.Do(pingCommand))
}

// Set sets a key/value pair
func (c *Client) Set(key string, value string) (string, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.String(connection.Do(setCommand, key, value))
}

// Get retrieves a key's value
func (c *Client) Get(key string) (string, bool, error) {
	connection := c.GetConnection()
	defer connection.Close()

	result, err := redis.String(connection.Do(getCommand, key))
	if err == redis.ErrNil {
		return result, false, nil
	}
	return result, true, err
}

// Del deletes keys
func (c *Client) Del(keys ...string) (int64, error) {
	connection := c.GetConnection()
	defer connection.Close()

	interfaces := make([]interface{}, len(keys))
	for i, key := range keys {
		interfaces[i] = key
	}
	return redis.Int64(connection.Do(delCommand, interfaces...))
}

// HSet sets a key's field/value pair
func (c *Client) HSet(key string, field string, value string) (int, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.Int(connection.Do(hsetCommand, key, field, value))
}

// HGet retrieves a key's field's value
func (c *Client) HGet(key string, field string) (string, bool, error) {
	connection := c.GetConnection()
	defer connection.Close()

	result, err := redis.String(connection.Do(hgetCommand, key, field))
	if err == redis.ErrNil {
		return result, false, nil
	}
	return result, true, err
}

// HGetAll retrieves the key
func (c *Client) HGetAll(key string) (map[string]string, error) {
	connection := c.GetConnection()
	defer connection.Close()

	results, err := redis.Strings(connection.Do(hgetAllCommand, key))
	if err != nil {
		return nil, err
	}

	resultsMap := make(map[string]string)
	for i := 0; i < len(results); i = i + 2 {
		key := results[i]
		value := results[i+1]
		resultsMap[key] = value
	}
	return resultsMap, err
}

// HDel deletes a key's fields
func (c *Client) HDel(key string, fields ...string) (int64, error) {
	connection := c.GetConnection()
	defer connection.Close()

	interfaces := make([]interface{}, len(fields)+1)
	interfaces[0] = key
	for i, key := range fields {
		interfaces[i+1] = key
	}
	return redis.Int64(connection.Do(hdelCommand, interfaces...))
}

// Incr increments the key's value
func (c *Client) Incr(key string) (int64, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.Int64(connection.Do(incrCommand, key))
}

// Decr decrements the key's value
func (c *Client) Decr(key string) (int64, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.Int64(connection.Do(decrCommand, key))
}

// IncrBy increments the key's value by the increment provided
func (c *Client) IncrBy(key string, increment int) (int64, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.Int64(connection.Do(incrByCommand, key, increment))
}

// DecrBy decrements the key's value by the decrement provided
func (c *Client) DecrBy(key string, decrement int) (int64, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.Int64(connection.Do(decrByCommand, key, decrement))
}

// Echo echoes the message
func (c *Client) Echo(message string) (string, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.String(connection.Do(echoCommand, message))
}

// Info returns redis information and statistics
func (c *Client) Info() (string, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.String(connection.Do(infoCommand))
}

// Close closes connections pool
func (c *Client) Close() error {
	return c.pool.Close()
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
