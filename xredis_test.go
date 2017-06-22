package xredis

import (
	"github.com/garyburd/redigo/redis"
	"testing"
)

func TestDefaultClient(t *testing.T) {
	client := DefaultClient()
	defer client.Close()
}

func TestSetupClient(t *testing.T) {
	client := SetupClient(&Options{})
	defer client.Close()
}

func TestNewClient(t *testing.T) {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	client := NewClient(pool)
	defer client.Close()
}

func TestClient_Close(t *testing.T) {
	client := DefaultClient()
	defer client.Close()
}

func TestClient_GetConnection(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	connection := client.GetConnection()
	defer connection.Close()
}

func TestClient_Ping(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.Ping()
}

func TestClient_Info(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.Info()
}

func TestClient_Echo(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.Echo("Hello")
}

func TestClient_FlushDb(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.FlushDb()
}

func TestClient_FlushAll(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.FlushAll()
}

func TestClient_Set(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.Set("key", "value")
	client.Del("key")
}

func TestClient_Get(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.Set("key", "value")
	client.Get("key")
	client.Get("unknown")
	client.Del("key")
}

func TestClient_Del(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.Del("key")
}

func TestClient_Keys(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.Keys("key")
}

func TestClient_Exists(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.Exists("key")
}

func TestClient_Incr(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.Incr("key")
	client.Del("key")
}

func TestClient_IncrBy(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.IncrBy("key", 1)
	client.Del("key")
}

func TestClient_IncrByFloat(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.IncrByFloat("key", 1.1)
	client.Del("key")
}

func TestClient_Decr(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.Decr("key")
	client.Del("key")
}

func TestClient_DecrBy(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.DecrBy("key", 1)
	client.Del("key")
}

func TestClient_DecrByFloat(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.DecrByFloat("key", 1.1)
	client.Del("key")
}

func TestClient_HSet(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.HSet("key", "field", "value")
	client.HDel("key", "field")
}

func TestClient_HGet(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.HSet("key", "field", "value")
	client.HGet("key", "field")
	client.HGet("key", "unknown")
	client.HDel("key", "field")
}

func TestClient_HGetAll(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.HSet("key", "field", "value")
	client.HGetAll("key")
	client.HGetAll("unknown")
	client.HDel("key", "field")
}

func TestClient_HKeys(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.HSet("key", "field", "value")
	client.HKeys("key")
	client.HDel("key", "field")
}

func TestClient_HExists(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.HSet("key", "field", "value")
	client.HExists("key", "field")
	client.HDel("key", "field")
}

func TestClient_HIncr(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.HIncr("key", "field")
	client.HDel("key", "field")
}

func TestClient_HIncrBy(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.HIncrBy("key", "field", 1)
	client.HDel("key", "field")
}

func TestClient_HIncrByFloat(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.HIncrByFloat("key", "field", 1.1)
	client.HDel("key", "field")
}

func TestClient_HDecr(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.HDecr("key", "field")
	client.HDel("key", "field")
}

func TestClient_HDecrBy(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.HDecrBy("key", "field", 1)
	client.HDel("key", "field")
}

func TestClient_HDecrByFloat(t *testing.T) {
	client := DefaultClient()
	defer client.Close()

	client.HDecrByFloat("key", "field", 1.1)
	client.HDel("key", "field")
}
