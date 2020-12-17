package command

import (
	"testing"

	"github.com/sampointer/digg/ranges"
	"github.com/stretchr/testify/require"
)

const ipv4 = "8.8.8.8"
const ipv6 = "2600:1901:0:ffff:ffff:ffff:ffff:aaaa"

func TestLookup(t *testing.T) {
	t.Parallel()
	t.Run("looks up IPv4 address", func(t *testing.T) {
		t.Parallel()
		prefix := ranges.Prefix{
			IPV4Prefix: "8.8.8.0/24",
		}

		p, err := Lookup(ipv4)
		require.NoError(t, err)
		require.Equal(t, 1, len(p))
		require.Equal(t, prefix, p[0])
	})

	t.Run("looks up IPv6 address", func(t *testing.T) {
		t.Parallel()
		prefix0 := ranges.Prefix{
			IPV6Prefix: "2600:1900::/28",
		}
		prefix1 := ranges.Prefix{
			IPV6Prefix: "2600:1901::/48",
			Scope:      "global",
			Service:    "Google Cloud",
		}

		p, err := Lookup(ipv6)
		require.NoError(t, err)
		require.Equal(t, 2, len(p))
		require.Equal(t, prefix0, p[0])
		require.Equal(t, prefix1, p[1])
	})
}
