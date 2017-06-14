package eredigo

import (
	"crypto/tls"
	"fmt"
	"time"
)

const (
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

	addressFormat = "%s:%d"
)

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

// GetAddress returns address
func (o *Options) GetAddress() string {
	return fmt.Sprintf(addressFormat, o.GetHost(), 6379)
}

// GetHost returns host
func (o *Options) GetHost() string {
	if len(o.host) == 0 {
		return defaultHost
	}
	return o.host
}

// GetPort returns port
func (o *Options) GetPort() int {
	if o.port <= 0 {
		return defaultPort
	}
	return o.port
}

// GetPassword returns password
func (o *Options) GetPassword() string {
	if len(o.password) == 0 {
		return defaultPassword
	}
	return o.password
}

// GetDatabase returns database
func (o *Options) GetDatabase() int {
	if o.database < 0 {
		return defaultDatabase
	}
	return o.database
}

// GetNetwork returns network
func (o *Options) GetNetwork() string {
	if len(o.network) == 0 {
		return defaultNetwork
	}
	return o.network
}

// GetConnectTimeout returns connect timeout
func (o *Options) GetConnectTimeout() time.Duration {
	if o.connectTimeout < 0 {
		return defaultConnectTimeout
	}
	return o.connectTimeout
}

// GetWriteTimeout returns write timeout
func (o *Options) GetWriteTimeout() time.Duration {
	if o.connectTimeout < 0 {
		return defaultWriteTimeout
	}
	return o.connectTimeout
}

// GetReadTimeout returns read timeout
func (o *Options) GetReadTimeout() time.Duration {
	if o.connectTimeout < 0 {
		return defaultReadTimeout
	}
	return o.connectTimeout
}

// GetConnectionIdleTimeout returns connection idle timeout
func (o *Options) GetConnectionIdleTimeout() time.Duration {
	if o.connectionIdleTimeout < 0 {
		return defaultConnectionIdleTimeout
	}
	return o.connectionIdleTimeout
}

// GetConnectionMaxIdle returns connection max idle
func (o *Options) GetConnectionMaxIdle() int {
	if o.connectionMaxIdle < 0 {
		return defaultConnectionMaxIdle
	}
	return o.connectionMaxIdle
}

// GetConnectionMaxActive returns connection max active
func (o *Options) GetConnectionMaxActive() int {
	if o.connectionMaxActive < 0 {
		return defaultConnectionMaxActive
	}
	return o.connectionMaxActive
}

// GetConnectionWait returns connection wait
func (o *Options) GetConnectionWait() bool {
	return o.connectionWait
}

// GetTlsConfig returns tls config
func (o *Options) GetTlsConfig() *tls.Config {
	return o.tlsConfig
}

// GetTlsSkipVerify returns tls skip verify
func (o *Options) GetTlsSkipVerify() bool {
	return o.tlsSkipVerify
}

// GetTestOnBorrowPeriod return test on borrow period
func (o *Options) GetTestOnBorrowPeriod() time.Duration {
	if o.testOnBorrowPeriod < 0 {
		return defaultTestOnBorrowTimeout
	}
	return o.testOnBorrowPeriod
}
