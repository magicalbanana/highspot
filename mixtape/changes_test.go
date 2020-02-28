package mixtape

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadChangesFromFile(t *testing.T) {
	t.Run("file exists", func(t *testing.T) {
		mixtape, err := LoadChangesFromFile("../testdata/changes.json")
		require.NoError(t, err)
		require.NotNil(t, mixtape)
	})

	t.Run("file does not exists", func(t *testing.T) {
		mixtape, err := LoadChangesFromFile("../testdata/doesnotexist.json")
		require.Error(t, err)
		require.Nil(t, mixtape)
	})
}
