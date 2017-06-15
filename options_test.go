package xredis

import (
	"testing"
	"github.com/stretchr/testify/assert"
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
