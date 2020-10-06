package proto

import (
	"net"
	"testing"

	"github.com/carrotsong/stun"
	"github.com/stretchr/testify/assert"
)

func TestRelayedAddress(t *testing.T) {
	// Simple tests because already tested in stun.
	a := RelayedAddress{
		IP:   net.IPv4(111, 11, 1, 2),
		Port: 333,
	}
	t.Run("String", func(t *testing.T) {
		if a.String() != "111.11.1.2:333" {
			t.Error("invalid string")
		}
	})
	m := new(stun.Message)
	if err := a.AddTo(m); err != nil {
		t.Fatal(err)
	}
	m.WriteHeader()
	decoded := new(stun.Message)

	_, err := decoded.Write(m.Raw)
	assert.NoError(t, err)

	var aGot RelayedAddress
	assert.NoError(t, aGot.GetFrom(decoded))
}
