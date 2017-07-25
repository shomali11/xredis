package xredis

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	expireOption    = "EX"
	notExistsOption = "NX"

	setCommand          = "SET"
	delCommand          = "DEL"
	getCommand          = "GET"
	keysCommand         = "KEYS"
	pingCommand         = "PING"
	echoCommand         = "ECHO"
	infoCommand         = "INFO"
	hSetCommand         = "HSET"
	hGetCommand         = "HGET"
	hDelCommand         = "HDEL"
	hKeysCommand        = "HKEYS"
	appendCommand       = "APPEND"
	expireCommand       = "EXPIRE"
	flushDbCommand      = "FLUSHDB"
	flushAllCommand     = "FLUSHALL"
	existsCommand       = "EXISTS"
	hExistsCommand      = "HEXISTS"
	hGetAllCommand      = "HGETALL"
	incrByCommand       = "INCRBY"
	incrByFloatCommand  = "INCRBYFLOAT"
	hIncrByCommand      = "HINCRBY"
	hIncrByFloatCommand = "HINCRBYFLOAT"
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

// FlushDb flushes the keys of the current database
func (c *Client) FlushDb() error {
	connection := c.GetConnection()
	defer connection.Close()

	return toError(connection.Do(flushDbCommand))
}

// FlushAll flushes the keys of all databases
func (c *Client) FlushAll() error {
	connection := c.GetConnection()
	defer connection.Close()

	return toError(connection.Do(flushAllCommand))
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

// Append to a key's value
func (c *Client) Append(key string, value string) (int64, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.Int64(connection.Do(appendCommand, key, value))
}

// Expire sets a key's timeout in seconds
func (c *Client) Expire(key string, timeout int) (bool, error) {
	connection := c.GetConnection()
	defer connection.Close()

	count, err := redis.Int64(connection.Do(expireCommand, key, timeout))
	return count > 0, err
}

// Set sets a key/value pair
func (c *Client) Set(key string, value string) (bool, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return toBool(connection.Do(setCommand, key, value))
}

// SetNx sets a key/value pair if the key does not exist
func (c *Client) SetNx(key string, value string) (bool, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return toBool(connection.Do(setCommand, key, value, notExistsOption))
}

// SetEx sets a key/value pair with a timeout in seconds
func (c *Client) SetEx(key string, value string, timeout int) (bool, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return toBool(connection.Do(setCommand, key, value, expireOption, timeout))
}

// Get retrieves a key's value
func (c *Client) Get(key string) (string, bool, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return toString(connection.Do(getCommand, key))
}

// Exists checks how many keys exist
func (c *Client) Exists(keys ...string) (bool, error) {
	connection := c.GetConnection()
	defer connection.Close()

	interfaces := make([]interface{}, len(keys))
	for i, key := range keys {
		interfaces[i] = key
	}
	count, err := redis.Int64(connection.Do(existsCommand, interfaces...))
	return count > 0, err
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

// Keys retrieves keys that match a pattern
func (c *Client) Keys(pattern string) ([]string, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.Strings(connection.Do(keysCommand, pattern))
}

// Incr increments the key's value
func (c *Client) Incr(key string) (int64, error) {
	return c.IncrBy(key, 1)
}

// IncrBy increments the key's value by the increment provided
func (c *Client) IncrBy(key string, increment int) (int64, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.Int64(connection.Do(incrByCommand, key, increment))
}

// IncrByFloat increments the key's value by the increment provided
func (c *Client) IncrByFloat(key string, increment float64) (float64, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.Float64(connection.Do(incrByFloatCommand, key, increment))
}

// Decr decrements the key's value
func (c *Client) Decr(key string) (int64, error) {
	return c.IncrBy(key, -1)
}

// DecrBy decrements the key's value by the decrement provided
func (c *Client) DecrBy(key string, decrement int) (int64, error) {
	return c.IncrBy(key, -decrement)
}

// DecrByFloat decrements the key's value by the decrement provided
func (c *Client) DecrByFloat(key string, decrement float64) (float64, error) {
	return c.IncrByFloat(key, -decrement)
}

// HSet sets a key's field/value pair
func (c *Client) HSet(key string, field string, value string) (int, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.Int(connection.Do(hSetCommand, key, field, value))
}

// HKeys retrieves a hash's keys
func (c *Client) HKeys(key string) ([]string, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.Strings(connection.Do(hKeysCommand, key))
}

// HExists determine's a key's field's existence
func (c *Client) HExists(key string, field string) (bool, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.Bool(connection.Do(hExistsCommand, key, field))
}

// HGet retrieves a key's field's value
func (c *Client) HGet(key string, field string) (string, bool, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return toString(connection.Do(hGetCommand, key, field))
}

// HGetAll retrieves the key
func (c *Client) HGetAll(key string) (map[string]string, error) {
	connection := c.GetConnection()
	defer connection.Close()

	results, err := redis.Strings(connection.Do(hGetAllCommand, key))
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
	return redis.Int64(connection.Do(hDelCommand, interfaces...))
}

// HIncr increments the key's field's value
func (c *Client) HIncr(key string, field string) (int64, error) {
	return c.HIncrBy(key, field, 1)
}

// HIncrBy increments the key's field's value by the increment provided
func (c *Client) HIncrBy(key string, field string, increment int) (int64, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.Int64(connection.Do(hIncrByCommand, key, field, increment))
}

// HIncrByFloat increments the key's field's value by the increment provided
func (c *Client) HIncrByFloat(key string, field string, increment float64) (float64, error) {
	connection := c.GetConnection()
	defer connection.Close()

	return redis.Float64(connection.Do(hIncrByFloatCommand, key, field, increment))
}

// HDecr decrements the key's field's value
func (c *Client) HDecr(key string, field string) (int64, error) {
	return c.HIncrBy(key, field, -1)
}

// HDecrBy decrements the key's field's value by the decrement provided
func (c *Client) HDecrBy(key string, field string, decrement int) (int64, error) {
	return c.HIncrBy(key, field, -decrement)
}

// HDecrByFloat decrements the key's field's value by the decrement provided
func (c *Client) HDecrByFloat(key string, field string, decrement float64) (float64, error) {
	return c.HIncrByFloat(key, field, -decrement)
}

// Close closes connections pool
func (c *Client) Close() error {
	return c.pool.Close()
}

func toError(reply interface{}, err error) error {
	_, _, e := toString(reply, err)
	return e
}

func toBool(reply interface{}, err error) (bool, error) {
	_, ok, e := toString(reply, err)
	return ok, e
}

func toString(reply interface{}, err error) (string, bool, error) {
	result, e := redis.String(reply, err)
	if e == redis.ErrNil {
		return result, false, nil
	}
	if e != nil {
		return result, false, e
	}
	return result, true, nil
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
