package ranges

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var prefixIPv4 = Prefix{
	IPV4Prefix: "34.80.0.0/15",
	Scope:      "asia-east1",
	Service:    "Google Cloud",
}

var prefixIPv6 = Prefix{
	IPV6Prefix: "2600:1901::/48",
	Scope:      "global",
	Service:    "Google Cloud",
}

var prefixBoth = Prefix{
	IPV4Prefix: "34.80.0.0/15",
	IPV6Prefix: "2600:1901::/48",
	Scope:      "global",
	Service:    "Google Cloud",
}

func TestPrefixString(t *testing.T) {
	t.Parallel()
	t.Run("IPv4", func(t *testing.T) {
		t.Parallel()
		require.Equal(
			t,
			"prefix: 34.80.0.0/15 scope: asia-east1 service: Google Cloud",
			prefixIPv4.String(),
		)
	})

	t.Run("IPv6", func(t *testing.T) {
		t.Parallel()
		require.Equal(
			t,
			"prefix: 2600:1901::/48 scope: global service: Google Cloud",
			prefixIPv6.String(),
		)
	})

	t.Run("both", func(t *testing.T) {
		t.Parallel()
		require.Equal(
			t,
			"prefix: 34.80.0.0/15,2600:1901::/48 scope: global service: Google Cloud",
			prefixBoth.String(),
		)
	})
}

func TestPrefixJSON(t *testing.T) {
	t.Parallel()
	t.Run("IPv4", func(t *testing.T) {
		t.Parallel()
		out, err := prefixIPv4.JSON()
		require.NoError(t, err)
		require.Equal(
			t,
			`{"ipv4Prefix":"34.80.0.0/15","scope":"asia-east1","service":"Google Cloud"}`,
			out,
		)
	})

	t.Run("IPv6", func(t *testing.T) {
		t.Parallel()
		out, err := prefixIPv6.JSON()
		require.NoError(t, err)
		require.Equal(
			t,
			`{"ipv6Prefix":"2600:1901::/48","scope":"global","service":"Google Cloud"}`,
			out,
		)
	})

	t.Run("both", func(t *testing.T) {
		t.Parallel()
		out, err := prefixBoth.JSON()
		require.NoError(t, err)
		require.Equal(
			t,
			`{"ipv4Prefix":"34.80.0.0/15","ipv6Prefix":"2600:1901::/48","scope":"global","service":"Google Cloud"}`,
			out,
		)
	})
}
