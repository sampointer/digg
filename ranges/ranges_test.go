package ranges

import (
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRanges(t *testing.T) {
	t.Parallel()
	doc, err := os.Open("../testdata/cloud.json")
	require.NoError(t, err)

	ranges, err := New(doc)
	require.NoError(t, err)

	t.Run("has IPv4 prefixes", func(t *testing.T) {
		t.Parallel()
		var count int64
		for _, p := range ranges.Prefixes {
			if p.IPV4Prefix != "" {
				count++
			}
		}
		require.NotZero(t, count, "should have 1 or more prefixes")
	})

	t.Run("has IPv6 prefixes", func(t *testing.T) {
		t.Parallel()
		var count int64
		for _, p := range ranges.Prefixes {
			if p.IPV6Prefix != "" {
				count++
			}
		}
		require.NotEmpty(t, count, "should have 1 or more prefixes")
	})

	t.Run("returns a Prefix struct for an IPv4 address", func(t *testing.T) {
		t.Parallel()
		prefix := Prefix{
			IPV4Prefix: "34.80.0.0/15",
			Service:    "Google Cloud",
			Scope:      "asia-east1",
		}

		ip := net.ParseIP("34.80.0.1")
		results, err := ranges.LookupIPv4(ip)
		require.NoError(t, err)
		require.Equal(t, prefix, results[0])
	})

	t.Run("returns no Prefix struct for non-Google IPv4 address", func(t *testing.T) {
		t.Parallel()
		ip := net.ParseIP("1.2.3.4")
		results, err := ranges.LookupIPv4(ip)
		require.NoError(t, err)
		require.Zero(t, len(results))
	})

	t.Run("returns a Prefix struct for an IPv6 address", func(t *testing.T) {
		t.Parallel()
		prefix := Prefix{
			IPV6Prefix: "2600:1901::/48",
			Service:    "Google Cloud",
			Scope:      "global",
		}

		ip := net.ParseIP("2600:1901:0:ffff:ffff:ffff:ffff:aaaa")
		results, err := ranges.LookupIPv6(ip)
		require.NoError(t, err)
		require.Equal(t, prefix, results[0])
	})

	t.Run("returns no Prefix struct for non-Google IPv6 address", func(t *testing.T) {
		t.Parallel()
		ip := net.ParseIP("1:2:3:4:5")
		results, err := ranges.LookupIPv6(ip)
		require.NoError(t, err)
		require.Zero(t, len(results))
	})
}
