package kong

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpstreamValid(T *testing.T) {

	assert := assert.New(T)

	upstream := &Upstream{}
	assert.Equal(false, upstream.Valid())
	upstream = &Upstream{
		Name: String("host.com"),
	}
	assert.Equal(true, upstream.Valid())
}

func TestUpstreamString(T *testing.T) {
	assert := assert.New(T)

	upstream := &Upstream{}
	assert.Equal("[ nil nil ]", upstream.String())

	upstream = &Upstream{
		Name: String("host.com"),
	}
	assert.Equal("[ nil host.com ]", upstream.String())
}
