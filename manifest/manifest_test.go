package manifest

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetManifest(t *testing.T) {
	t.Parallel()

	t.Run("manifests are valid JSON", func(t *testing.T) {
		for _, doc := range GetManifest() {
			res, err := ioutil.ReadAll(doc)
			require.NoError(t, err)
			require.True(t, json.Valid(res))
		}
	})
}
