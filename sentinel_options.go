package xredis

import (
	"crypto/tls"
	"errors"
	"github.com/FZambia/go-sentinel"
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	masterRole          = "master"
	connectionRoleError = "connection failed master role check"

	defaultSentinelAddress               = "localhost:26379"
	defaultSentinelMasterName            = "master"
	defaultSentinelPassword              = ""
	defaultSentinelDatabase              = 0
	defaultSentinelNetwork               = "tcp"
	defaultSentinelConnectTimeout        = time.Second
	defaultSentinelWriteTimeout          = time.Second
	defaultSentinelReadTimeout           = time.Second
	defaultSentinelConnectionIdleTimeout = 240 * time.Second
	defaultSentinelConnectionMaxIdle     = 100
	defaultSentinelConnectionMaxActive   = 10000
)

// SentinelOptions contains redis sentinel options
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

// GetAddresses returns sentinel address
func (o *SentinelOptions) GetAddresses() []string {
	if len(o.Addresses) == 0 {
		return []string{defaultSentinelAddress}
	}
	return o.Addresses
}

// GetMasterName returns master name
func (o *SentinelOptions) GetMasterName() string {
	if len(o.MasterName) == 0 {
		return defaultSentinelMasterName
	}
	return o.MasterName
}

// GetPassword returns password
func (o *SentinelOptions) GetPassword() string {
	if len(o.Password) == 0 {
		return defaultSentinelPassword
	}
	return o.Password
}

// GetDatabase returns database
func (o *SentinelOptions) GetDatabase() int {
	if o.Database < 0 {
		return defaultSentinelDatabase
	}
	return o.Database
}

// GetNetwork returns network
func (o *SentinelOptions) GetNetwork() string {
	if len(o.Network) == 0 {
		return defaultSentinelNetwork
	}
	return o.Network
}

// GetConnectTimeout returns connect timeout
func (o *SentinelOptions) GetConnectTimeout() time.Duration {
	if o.ConnectTimeout < 0 {
		return defaultSentinelConnectTimeout
	}
	return o.ConnectTimeout
}

// GetWriteTimeout returns write timeout
func (o *SentinelOptions) GetWriteTimeout() time.Duration {
	if o.WriteTimeout < 0 {
		return defaultSentinelWriteTimeout
	}
	return o.WriteTimeout
}

// GetReadTimeout returns read timeout
func (o *SentinelOptions) GetReadTimeout() time.Duration {
	if o.ReadTimeout < 0 {
		return defaultSentinelReadTimeout
	}
	return o.ReadTimeout
}

// GetConnectionIdleTimeout returns connection idle timeout
func (o *SentinelOptions) GetConnectionIdleTimeout() time.Duration {
	if o.ConnectionIdleTimeout < 0 {
		return defaultSentinelConnectionIdleTimeout
	}
	return o.ConnectionIdleTimeout
}

// GetConnectionMaxIdle returns connection max idle
func (o *SentinelOptions) GetConnectionMaxIdle() int {
	if o.ConnectionMaxIdle < 0 {
		return defaultSentinelConnectionMaxIdle
	}
	return o.ConnectionMaxIdle
}

// GetConnectionMaxActive returns connection max active
func (o *SentinelOptions) GetConnectionMaxActive() int {
	if o.ConnectionMaxActive < 0 {
		return defaultSentinelConnectionMaxActive
	}
	return o.ConnectionMaxActive
}

// GetConnectionWait returns connection wait
func (o *SentinelOptions) GetConnectionWait() bool {
	return o.ConnectionWait
}

// GetTlsConfig returns tls config
func (o *SentinelOptions) GetTlsConfig() *tls.Config {
	return o.TlsConfig
}

// GetTlsSkipVerify returns tls skip verify
func (o *SentinelOptions) GetTlsSkipVerify() bool {
	return o.TlsSkipVerify
}

func newSentinelPool(options *SentinelOptions) *redis.Pool {
	connectionIdleTimeout := options.GetConnectionIdleTimeout()
	connectionMaxActive := options.GetConnectionMaxActive()
	connectionMaxIdle := options.GetConnectionMaxIdle()
	connectionWait := options.GetConnectionWait()

	return &redis.Pool{
		IdleTimeout:  connectionIdleTimeout,
		MaxActive:    connectionMaxActive,
		MaxIdle:      connectionMaxIdle,
		Wait:         connectionWait,
		Dial:         sentinelDial(options),
		TestOnBorrow: sentinelTestOnBorrow(),
	}
}

func sentinelDial(options *SentinelOptions) func() (redis.Conn, error) {
	sentinelNetwork := options.GetNetwork()

	dialSentinelOptions := make([]redis.DialOption, 5)
	dialSentinelOptions[0] = redis.DialConnectTimeout(options.GetConnectTimeout())
	dialSentinelOptions[1] = redis.DialWriteTimeout(options.GetWriteTimeout())
	dialSentinelOptions[2] = redis.DialReadTimeout(options.GetReadTimeout())
	dialSentinelOptions[3] = redis.DialTLSSkipVerify(options.GetTlsSkipVerify())
	dialSentinelOptions[4] = redis.DialTLSConfig(options.GetTlsConfig())

	sentinelDetails := &sentinel.Sentinel{
		Addrs:      options.GetAddresses(),
		MasterName: options.GetMasterName(),
		Dial: func(address string) (redis.Conn, error) {
			connection, err := redis.Dial(sentinelNetwork, address, dialSentinelOptions...)
			if err != nil {
				return nil, err
			}
			return connection, nil
		},
	}

	network := options.GetNetwork()

	dialServerOptions := make([]redis.DialOption, 7)
	dialServerOptions[0] = redis.DialPassword(options.GetPassword())
	dialServerOptions[1] = redis.DialDatabase(options.GetDatabase())
	dialServerOptions[2] = redis.DialConnectTimeout(options.GetConnectTimeout())
	dialServerOptions[3] = redis.DialWriteTimeout(options.GetWriteTimeout())
	dialServerOptions[4] = redis.DialReadTimeout(options.GetReadTimeout())
	dialServerOptions[5] = redis.DialTLSSkipVerify(options.GetTlsSkipVerify())
	dialServerOptions[6] = redis.DialTLSConfig(options.GetTlsConfig())

	return func() (redis.Conn, error) {
		address, err := sentinelDetails.MasterAddr()
		if err != nil {
			return nil, err
		}

		connection, err := redis.Dial(network, address, dialServerOptions...)
		if err != nil {
			return nil, err
		}
		return connection, nil
	}
}

func sentinelTestOnBorrow() func(redis.Conn, time.Time) error {
	return func(connection redis.Conn, t time.Time) error {
		if !sentinel.TestRole(connection, masterRole) {
			return errors.New(connectionRoleError)
		} else {
			return nil
		}
	}
}
