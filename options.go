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
	host                  string
	port                  int
	password              string
	database              int
	network               string
	connectTimeout        time.Duration
	writeTimeout          time.Duration
	readTimeout           time.Duration
	connectionIdleTimeout time.Duration
	connectionMaxIdle     int
	connectionMaxActive   int
	connectionWait        bool
	tlsConfig             *tls.Config
	tlsSkipVerify         bool
	testOnBorrowPeriod    time.Duration
}

// Address returns address
func (o *Options) Address() string {
	return fmt.Sprintf(addressFormat, o.Host(), 6379)
}

// Host returns host
func (o *Options) Host() string {
	if len(o.host) == 0 {
		return defaultHost
	}
	return o.host
}

// Port returns port
func (o *Options) Port() int {
	if o.port <= 0 {
		return defaultPort
	}
	return o.port
}

// Password returns password
func (o *Options) Password() string {
	if len(o.password) == 0 {
		return defaultPassword
	}
	return o.password
}

// Database returns database
func (o *Options) Database() int {
	if o.database < 0 {
		return defaultDatabase
	}
	return o.database
}

// Network returns network
func (o *Options) Network() string {
	if len(o.network) == 0 {
		return defaultNetwork
	}
	return o.network
}

// ConnectTimeout returns connect timeout
func (o *Options) ConnectTimeout() time.Duration {
	if o.connectTimeout < 0 {
		return defaultConnectTimeout
	}
	return o.connectTimeout
}

// WriteTimeout returns write timeout
func (o *Options) WriteTimeout() time.Duration {
	if o.connectTimeout < 0 {
		return defaultWriteTimeout
	}
	return o.connectTimeout
}

// ReadTimeout returns read timeout
func (o *Options) ReadTimeout() time.Duration {
	if o.connectTimeout < 0 {
		return defaultReadTimeout
	}
	return o.connectTimeout
}

// ConnectionIdleTimeout returns connection idle timeout
func (o *Options) ConnectionIdleTimeout() time.Duration {
	if o.connectionIdleTimeout < 0 {
		return defaultConnectionIdleTimeout
	}
	return o.connectionIdleTimeout
}

// ConnectionMaxIdle returns connection max idle
func (o *Options) ConnectionMaxIdle() int {
	if o.connectionMaxIdle < 0 {
		return defaultConnectionMaxIdle
	}
	return o.connectionMaxIdle
}

// ConnectionMaxActive returns connection max active
func (o *Options) ConnectionMaxActive() int {
	if o.connectionMaxActive < 0 {
		return defaultConnectionMaxActive
	}
	return o.connectionMaxActive
}

// ConnectionWait returns connection wait
func (o *Options) ConnectionWait() bool {
	return o.connectionWait
}

// TlsConfig returns tls config
func (o *Options) TlsConfig() *tls.Config {
	return o.tlsConfig
}

// TlsSkipVerify returns tls skip verify
func (o *Options) TlsSkipVerify() bool {
	return o.tlsSkipVerify
}

// TestOnBorrowPeriod return test on borrow period
func (o *Options) TestOnBorrowPeriod() time.Duration {
	if o.testOnBorrowPeriod < 0 {
		return defaultTestOnBorrowTimeout
	}
	return o.testOnBorrowPeriod
}
