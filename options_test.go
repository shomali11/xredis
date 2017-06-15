package xredis

import (
	"crypto/tls"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOptions_GetAddress(t *testing.T) {
	options := Options{}
	assert.Equal(t, options.GetAddress(), "localhost:6379")

	options = Options{Host: "abc"}
	assert.Equal(t, options.GetAddress(), "abc:6379")

	options = Options{Port: 1}
	assert.Equal(t, options.GetAddress(), "localhost:1")

	options = Options{Port: 0}
	assert.Equal(t, options.GetAddress(), "localhost:6379")
}

func TestOptions_GetPassword(t *testing.T) {
	options := Options{}
	assert.Equal(t, options.GetPassword(), defaultPassword)

	options = Options{Password: "abc"}
	assert.Equal(t, options.GetPassword(), "abc")
}

func TestOptions_GetDatabase(t *testing.T) {
	options := Options{Database: 1}
	assert.Equal(t, options.GetDatabase(), 1)

	options = Options{Database: 0}
	assert.Equal(t, options.GetDatabase(), 0)

	options = Options{Database: -1}
	assert.Equal(t, options.GetDatabase(), defaultDatabase)
}

func TestOptions_GetNetwork(t *testing.T) {
	options := Options{}
	assert.Equal(t, options.GetNetwork(), defaultNetwork)

	options = Options{Network: "abc"}
	assert.Equal(t, options.GetNetwork(), "abc")
}

func TestOptions_GetConnectTimeout(t *testing.T) {
	options := Options{ConnectTimeout: 1}
	assert.Equal(t, options.GetConnectTimeout(), time.Duration(1))

	options = Options{ConnectTimeout: 0}
	assert.Equal(t, options.GetConnectTimeout(), time.Duration(0))

	options = Options{ConnectTimeout: -1}
	assert.Equal(t, options.GetConnectTimeout(), defaultConnectTimeout)
}

func TestOptions_GetWriteTimeout(t *testing.T) {
	options := Options{WriteTimeout: 1}
	assert.Equal(t, options.GetWriteTimeout(), time.Duration(1))

	options = Options{WriteTimeout: 0}
	assert.Equal(t, options.GetWriteTimeout(), time.Duration(0))

	options = Options{WriteTimeout: -1}
	assert.Equal(t, options.GetWriteTimeout(), defaultWriteTimeout)
}

func TestOptions_GetReadTimeout(t *testing.T) {
	options := Options{ReadTimeout: 1}
	assert.Equal(t, options.GetReadTimeout(), time.Duration(1))

	options = Options{ReadTimeout: 0}
	assert.Equal(t, options.GetReadTimeout(), time.Duration(0))

	options = Options{ReadTimeout: -1}
	assert.Equal(t, options.GetReadTimeout(), defaultReadTimeout)
}

func TestOptions_GetConnectionIdleTimeout(t *testing.T) {
	options := Options{ConnectionIdleTimeout: 1}
	assert.Equal(t, options.GetConnectionIdleTimeout(), time.Duration(1))

	options = Options{ConnectionIdleTimeout: 0}
	assert.Equal(t, options.GetConnectionIdleTimeout(), time.Duration(0))

	options = Options{ConnectionIdleTimeout: -1}
	assert.Equal(t, options.GetConnectionIdleTimeout(), defaultConnectionIdleTimeout)
}

func TestOptions_GetConnectionMaxIdle(t *testing.T) {
	options := Options{ConnectionMaxIdle: 1}
	assert.Equal(t, options.GetConnectionMaxIdle(), 1)

	options = Options{ConnectionMaxIdle: 0}
	assert.Equal(t, options.GetConnectionMaxIdle(), 0)

	options = Options{ConnectionMaxIdle: -1}
	assert.Equal(t, options.GetConnectionMaxIdle(), defaultConnectionMaxIdle)
}

func TestOptions_GetConnectionMaxActive(t *testing.T) {
	options := Options{ConnectionMaxActive: 1}
	assert.Equal(t, options.GetConnectionMaxActive(), 1)

	options = Options{ConnectionMaxActive: 0}
	assert.Equal(t, options.GetConnectionMaxActive(), 0)

	options = Options{ConnectionMaxActive: -1}
	assert.Equal(t, options.GetConnectionMaxActive(), defaultConnectionMaxActive)
}

func TestOptions_GetConnectionWait(t *testing.T) {
	options := Options{ConnectionWait: true}
	assert.Equal(t, options.GetConnectionWait(), true)

	options = Options{ConnectionWait: false}
	assert.Equal(t, options.GetConnectionWait(), false)
}

func TestOptions_GetTlsConfig(t *testing.T) {
	options := Options{}
	assert.Nil(t, options.GetTlsConfig())

	config := &tls.Config{}
	options = Options{TlsConfig: config}
	assert.Equal(t, options.GetTlsConfig(), config)
}

func TestOptions_GetTlsSkipVerify(t *testing.T) {
	options := Options{TlsSkipVerify: true}
	assert.Equal(t, options.GetTlsSkipVerify(), true)

	options = Options{TlsSkipVerify: false}
	assert.Equal(t, options.GetTlsSkipVerify(), false)
}

func TestOptions_GetTestOnBorrowPeriod(t *testing.T) {
	options := Options{TestOnBorrowPeriod: 1}
	assert.Equal(t, options.GetTestOnBorrowPeriod(), time.Duration(1))

	options = Options{TestOnBorrowPeriod: 0}
	assert.Equal(t, options.GetTestOnBorrowPeriod(), time.Duration(0))

	options = Options{TestOnBorrowPeriod: -1}
	assert.Equal(t, options.GetTestOnBorrowPeriod(), defaultTestOnBorrowTimeout)
}
