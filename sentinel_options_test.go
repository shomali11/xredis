package xredis

import (
	"crypto/tls"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSentinelOptions_GetAddresses(t *testing.T) {
	options := SentinelOptions{}
	assert.Equal(t, options.GetAddresses(), []string{defaultSentinelAddress})

	options = SentinelOptions{Addresses: []string{"a:1"}}
	assert.Equal(t, options.GetAddresses(), []string{"a:1"})
}

func TestSentinelOptions_GetMasterName(t *testing.T) {
	options := SentinelOptions{}
	assert.Equal(t, options.GetMasterName(), defaultSentinelMasterName)

	options = SentinelOptions{MasterName: "abc"}
	assert.Equal(t, options.GetMasterName(), "abc")
}

func TestSentinelOptions_GetPassword(t *testing.T) {
	options := SentinelOptions{}
	assert.Equal(t, options.GetPassword(), defaultPassword)

	options = SentinelOptions{Password: "abc"}
	assert.Equal(t, options.GetPassword(), "abc")
}

func TestSentinelOptions_GetDatabase(t *testing.T) {
	options := SentinelOptions{Database: 1}
	assert.Equal(t, options.GetDatabase(), 1)

	options = SentinelOptions{Database: 0}
	assert.Equal(t, options.GetDatabase(), 0)

	options = SentinelOptions{Database: -1}
	assert.Equal(t, options.GetDatabase(), defaultDatabase)
}

func TestSentinelOptions_GetNetwork(t *testing.T) {
	options := SentinelOptions{}
	assert.Equal(t, options.GetNetwork(), defaultNetwork)

	options = SentinelOptions{Network: "abc"}
	assert.Equal(t, options.GetNetwork(), "abc")
}

func TestSentinelOptions_GetConnectTimeout(t *testing.T) {
	options := SentinelOptions{ConnectTimeout: 1}
	assert.Equal(t, options.GetConnectTimeout(), time.Duration(1))

	options = SentinelOptions{ConnectTimeout: 0}
	assert.Equal(t, options.GetConnectTimeout(), time.Duration(0))

	options = SentinelOptions{ConnectTimeout: -1}
	assert.Equal(t, options.GetConnectTimeout(), defaultConnectTimeout)
}

func TestSentinelOptions_GetWriteTimeout(t *testing.T) {
	options := SentinelOptions{WriteTimeout: 1}
	assert.Equal(t, options.GetWriteTimeout(), time.Duration(1))

	options = SentinelOptions{WriteTimeout: 0}
	assert.Equal(t, options.GetWriteTimeout(), time.Duration(0))

	options = SentinelOptions{WriteTimeout: -1}
	assert.Equal(t, options.GetWriteTimeout(), defaultWriteTimeout)
}

func TestSentinelOptions_GetReadTimeout(t *testing.T) {
	options := SentinelOptions{ReadTimeout: 1}
	assert.Equal(t, options.GetReadTimeout(), time.Duration(1))

	options = SentinelOptions{ReadTimeout: 0}
	assert.Equal(t, options.GetReadTimeout(), time.Duration(0))

	options = SentinelOptions{ReadTimeout: -1}
	assert.Equal(t, options.GetReadTimeout(), defaultReadTimeout)
}

func TestSentinelOptions_GetConnectionIdleTimeout(t *testing.T) {
	options := SentinelOptions{ConnectionIdleTimeout: 1}
	assert.Equal(t, options.GetConnectionIdleTimeout(), time.Duration(1))

	options = SentinelOptions{ConnectionIdleTimeout: 0}
	assert.Equal(t, options.GetConnectionIdleTimeout(), time.Duration(0))

	options = SentinelOptions{ConnectionIdleTimeout: -1}
	assert.Equal(t, options.GetConnectionIdleTimeout(), defaultConnectionIdleTimeout)
}

func TestSentinelOptions_GetConnectionMaxIdle(t *testing.T) {
	options := SentinelOptions{ConnectionMaxIdle: 1}
	assert.Equal(t, options.GetConnectionMaxIdle(), 1)

	options = SentinelOptions{ConnectionMaxIdle: 0}
	assert.Equal(t, options.GetConnectionMaxIdle(), 0)

	options = SentinelOptions{ConnectionMaxIdle: -1}
	assert.Equal(t, options.GetConnectionMaxIdle(), defaultConnectionMaxIdle)
}

func TestSentinelOptions_GetConnectionMaxActive(t *testing.T) {
	options := SentinelOptions{ConnectionMaxActive: 1}
	assert.Equal(t, options.GetConnectionMaxActive(), 1)

	options = SentinelOptions{ConnectionMaxActive: 0}
	assert.Equal(t, options.GetConnectionMaxActive(), 0)

	options = SentinelOptions{ConnectionMaxActive: -1}
	assert.Equal(t, options.GetConnectionMaxActive(), defaultConnectionMaxActive)
}

func TestSentinelOptions_GetConnectionWait(t *testing.T) {
	options := SentinelOptions{ConnectionWait: true}
	assert.Equal(t, options.GetConnectionWait(), true)

	options = SentinelOptions{ConnectionWait: false}
	assert.Equal(t, options.GetConnectionWait(), false)
}

func TestSentinelOptions_GetTlsConfig(t *testing.T) {
	options := SentinelOptions{}
	assert.Nil(t, options.GetTlsConfig())

	config := &tls.Config{}
	options = SentinelOptions{TlsConfig: config}
	assert.Equal(t, options.GetTlsConfig(), config)
}

func TestSentinelOptions_GetTlsSkipVerify(t *testing.T) {
	options := SentinelOptions{TlsSkipVerify: true}
	assert.Equal(t, options.GetTlsSkipVerify(), true)

	options = SentinelOptions{TlsSkipVerify: false}
	assert.Equal(t, options.GetTlsSkipVerify(), false)
}

func TestSentinelOptions_GetTestOnBorrowPeriod(t *testing.T) {
	options := SentinelOptions{TestOnBorrowPeriod: 1}
	assert.Equal(t, options.GetTestOnBorrowPeriod(), time.Duration(1))

	options = SentinelOptions{TestOnBorrowPeriod: 0}
	assert.Equal(t, options.GetTestOnBorrowPeriod(), time.Duration(0))

	options = SentinelOptions{TestOnBorrowPeriod: -1}
	assert.Equal(t, options.GetTestOnBorrowPeriod(), defaultTestOnBorrowTimeout)
}
