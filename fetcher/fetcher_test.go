package fetcher

import (
	"io/ioutil"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetcher(t *testing.T) {
	f, err := Fetch()
	require.NoError(t, err)
	require.Equal(t, 2, len(f))

	for _, doc := range f {
		d, err := ioutil.ReadAll(doc)
		require.NoError(t, err)

		re := regexp.MustCompile("syncToken")
		res := re.FindString(string(d))
		require.NotZero(t, res)
	}
}
