package xredis

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Close(t *testing.T) {
	connection := redigomock.NewConn()

	client := mockClient(connection)
	assert.Nil(t, client.Close())
}

func TestClient_Ping(t *testing.T) {
	connection := redigomock.NewConn()
	connection.Command("PING").Expect("PONG")

	client := mockClient(connection)

	result, err := client.Ping()
	assert.Equal(t, result, "PONG")
	assert.Nil(t, err)
}

func TestClient_Info(t *testing.T) {
	connection := redigomock.NewConn()
	connection.Command("INFO").Expect("INFO")

	client := mockClient(connection)

	result, err := client.Info()
	assert.Equal(t, result, "INFO")
	assert.Nil(t, err)
}

func TestClient_Echo(t *testing.T) {
	connection := redigomock.NewConn()
	connection.Command("ECHO", "Hello").Expect("Hello")

	client := mockClient(connection)

	result, err := client.Echo("Hello")
	assert.Equal(t, result, "Hello")
	assert.Nil(t, err)
}

func TestClient_Append(t *testing.T) {
	connection := redigomock.NewConn()
	connection.Command("APPEND", "name", "a").Expect(int64(1))

	client := mockClient(connection)

	number, err := client.Append("name", "a")
	assert.Equal(t, number, int64(1))
	assert.Nil(t, err)

	connection.Command("APPEND", "name", "b").Expect(int64(2))

	client = mockClient(connection)

	number, err = client.Append("name", "b")
	assert.Equal(t, number, int64(2))
	assert.Nil(t, err)
}

func TestClient_Expire(t *testing.T) {
	connection := redigomock.NewConn()
	connection.Command("EXPIRE", "name", 10).Expect(int64(1))

	client := mockClient(connection)

	ok, err := client.Expire("name", 10)
	assert.True(t, ok)
	assert.Nil(t, err)

	connection.Command("EXPIRE", "unknown", 10).Expect(int64(0))

	client = mockClient(connection)

	ok, err = client.Expire("unknown", 10)
	assert.False(t, ok)
	assert.Nil(t, err)
}

func TestClient_FlushDb(t *testing.T) {
	connection := redigomock.NewConn()
	connection.Command("FLUSHDB").Expect("OK")

	client := mockClient(connection)

	err := client.FlushDb()
	assert.Nil(t, err)
}

func TestClient_FlushAll(t *testing.T) {
	connection := redigomock.NewConn()
	connection.Command("FLUSHALL").Expect("OK")

	client := mockClient(connection)

	err := client.FlushAll()
	assert.Nil(t, err)
}

func TestClient_Set(t *testing.T) {
	connection := redigomock.NewConn()
	connection.Command("SET", "key", "value").Expect("OK")

	client := mockClient(connection)

	ok, err := client.Set("key", "value")
	assert.True(t, ok)
	assert.Nil(t, err)
}

func TestClient_SetNx(t *testing.T) {
	connection := redigomock.NewConn()
	connection.Command("SET", "key", "value", "NX").Expect("OK")

	client := mockClient(connection)

	ok, err := client.SetNx("key", "value")
	assert.True(t, ok)
	assert.Nil(t, err)

	connection.Command("SET", "key", "value", "NX").ExpectError(redis.ErrNil)

	client = mockClient(connection)

	ok, err = client.SetNx("key", "value")
	assert.False(t, ok)
	assert.Nil(t, err)
}

func TestClient_SetEx(t *testing.T) {
	connection := redigomock.NewConn()
	connection.Command("SET", "key", "value", "EX", 1).Expect("OK")

	client := mockClient(connection)

	ok, err := client.SetEx("key", "value", 1)
	assert.True(t, ok)
	assert.Nil(t, err)

	connection.Command("SET", "key", "value", "EX", 1).ExpectError(errors.New("Opps"))

	client = mockClient(connection)

	ok, err = client.SetEx("key", "value", 1)
	assert.False(t, ok)
	assert.NotNil(t, err)
}

func TestClient_Get(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("GET", "key").Expect("value")
	client := mockClient(connection)

	result, ok, err := client.Get("key")
	assert.Equal(t, result, "value")
	assert.True(t, ok)
	assert.Nil(t, err)

	connection.Command("GET", "unknown").ExpectError(redis.ErrNil)
	client = mockClient(connection)

	result, ok, err = client.Get("unknown")
	assert.Equal(t, result, "")
	assert.False(t, ok)
	assert.Nil(t, err)

	connection.Command("GET", "unknown").ExpectError(errors.New("Oops"))
	client = mockClient(connection)

	result, ok, err = client.Get("unknown")
	assert.Equal(t, result, "")
	assert.False(t, ok)
	assert.NotNil(t, err)
}

func TestClient_Del(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("DEL", "key").Expect(int64(1))
	client := mockClient(connection)

	result, err := client.Del("key")
	assert.Equal(t, result, int64(1))
	assert.Nil(t, err)
}

func TestClient_Keys(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("KEYS", "key").Expect([]interface{}{[]byte("key")})
	client := mockClient(connection)

	results, err := client.Keys("key")
	assert.Equal(t, len(results), 1)
	assert.Equal(t, results[0], "key")
	assert.Nil(t, err)
}

func TestClient_Exists(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("EXISTS", "key").Expect(int64(1))
	client := mockClient(connection)

	result, err := client.Exists("key")
	assert.True(t, result)
	assert.Nil(t, err)
}

func TestClient_Incr(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("INCRBY", "key", 1).Expect(int64(1))
	client := mockClient(connection)

	result, err := client.Incr("key")
	assert.Equal(t, result, int64(1))
	assert.Nil(t, err)
}

func TestClient_IncrBy(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("INCRBY", "key", 10).Expect(int64(10))
	client := mockClient(connection)

	result, err := client.IncrBy("key", 10)
	assert.Equal(t, result, int64(10))
	assert.Nil(t, err)
}

func TestClient_IncrByFloat(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("INCRBYFLOAT", "key", 5.5).Expect([]byte("5.5"))
	client := mockClient(connection)

	result, err := client.IncrByFloat("key", 5.5)
	assert.Equal(t, result, float64(5.5))
	assert.Nil(t, err)
}

func TestClient_Decr(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("INCRBY", "key", -1).Expect(int64(-1))
	client := mockClient(connection)

	result, err := client.Decr("key")
	assert.Equal(t, result, int64(-1))
	assert.Nil(t, err)
}

func TestClient_DecrBy(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("INCRBY", "key", -10).Expect(int64(-10))
	client := mockClient(connection)

	result, err := client.DecrBy("key", 10)
	assert.Equal(t, result, int64(-10))
	assert.Nil(t, err)
}

func TestClient_DecrByFloat(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("INCRBYFLOAT", "key", -5.5).Expect([]byte("-5.5"))
	client := mockClient(connection)

	result, err := client.DecrByFloat("key", 5.5)
	assert.Equal(t, result, float64(-5.5))
	assert.Nil(t, err)
}

func TestClient_HSet(t *testing.T) {
	connection := redigomock.NewConn()
	connection.Command("HSET", "key", "field", "value").Expect([]byte("1"))

	client := mockClient(connection)

	result, err := client.HSet("key", "field", "value")
	assert.Equal(t, result, int(1))
	assert.Nil(t, err)
}

func TestClient_HGet(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("HGET", "key", "field").Expect("value")
	client := mockClient(connection)

	result, ok, err := client.HGet("key", "field")
	assert.Equal(t, result, "value")
	assert.True(t, ok)
	assert.Nil(t, err)

	connection.Command("HGET", "unknown", "field").ExpectError(redis.ErrNil)
	client = mockClient(connection)

	result, ok, err = client.HGet("unknown", "field")
	assert.Equal(t, result, "")
	assert.False(t, ok)
	assert.Nil(t, err)

	connection.Command("HGET", "unknown", "field").ExpectError(errors.New("Oops"))
	client = mockClient(connection)

	result, ok, err = client.HGet("unknown", "field")
	assert.Equal(t, result, "")
	assert.False(t, ok)
	assert.NotNil(t, err)
}

func TestClient_HGetAll(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("HGETALL", "key").ExpectMap(map[string]string{"key": "value"})

	client := mockClient(connection)

	result, err := client.HGetAll("key")
	assert.Equal(t, result["key"], "value")
	assert.Nil(t, err)

	connection.Command("HGETALL", "unknown").ExpectError(errors.New("Oops"))
	client = mockClient(connection)

	result, err = client.HGetAll("unknown")
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestClient_HDel(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("HDEL", "key", "field").Expect(int64(1))
	client := mockClient(connection)

	result, err := client.HDel("key", "field")
	assert.Equal(t, result, int64(1))
	assert.Nil(t, err)
}

func TestClient_HKeys(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("HKEYS", "key").Expect([]interface{}{[]byte("key")})
	client := mockClient(connection)

	results, err := client.HKeys("key")
	assert.Equal(t, len(results), 1)
	assert.Equal(t, results[0], "key")
	assert.Nil(t, err)
}

func TestClient_HExists(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("HEXISTS", "key", "field").Expect(int64(1))
	client := mockClient(connection)

	result, err := client.HExists("key", "field")
	assert.True(t, result)
	assert.Nil(t, err)
}

func TestClient_HIncr(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("HINCRBY", "key", "field", 1).Expect(int64(1))
	client := mockClient(connection)

	result, err := client.HIncr("key", "field")
	assert.Equal(t, result, int64(1))
	assert.Nil(t, err)
}

func TestClient_HIncrBy(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("HINCRBY", "key", "field", 10).Expect(int64(10))
	client := mockClient(connection)

	result, err := client.HIncrBy("key", "field", 10)
	assert.Equal(t, result, int64(10))
	assert.Nil(t, err)
}

func TestClient_HIncrByFloat(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("HINCRBYFLOAT", "key", "field", 5.5).Expect([]byte("5.5"))
	client := mockClient(connection)

	result, err := client.HIncrByFloat("key", "field", 5.5)
	assert.Equal(t, result, float64(5.5))
	assert.Nil(t, err)
}

func TestClient_HDecr(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("HINCRBY", "key", "field", -1).Expect(int64(-1))
	client := mockClient(connection)

	result, err := client.HDecr("key", "field")
	assert.Equal(t, result, int64(-1))
	assert.Nil(t, err)
}

func TestClient_HDecrBy(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("HINCRBY", "key", "field", -10).Expect(int64(-10))
	client := mockClient(connection)

	result, err := client.HDecrBy("key", "field", 10)
	assert.Equal(t, result, int64(-10))
	assert.Nil(t, err)
}

func TestClient_HDecrByFloat(t *testing.T) {
	connection := redigomock.NewConn()

	connection.Command("HINCRBYFLOAT", "key", "field", -5.5).Expect([]byte("-5.5"))
	client := mockClient(connection)

	result, err := client.HDecrByFloat("key", "field", 5.5)
	assert.Equal(t, result, float64(-5.5))
	assert.Nil(t, err)
}

func TestDefaultClient(t *testing.T) {
	client := DefaultClient()
	defer client.Close()
}

func TestSetupClient(t *testing.T) {
	client := SetupClient(&Options{})
	defer client.Close()
}

func mockClient(connection *redigomock.Conn) *Client {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return connection, nil
		},
	}
	return NewClient(pool)
}
