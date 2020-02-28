package mixtape

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadMixtapeFromFile(t *testing.T) {
	t.Run("file exists", func(t *testing.T) {
		mixtape, err := LoadMixtapeFromFile("../testdata/mixtape.json")
		require.NoError(t, err)
		require.NotNil(t, mixtape)
	})

	t.Run("file does not exists", func(t *testing.T) {
		mixtape, err := LoadMixtapeFromFile("../testdata/doesnotexist.json")
		require.Error(t, err)
		require.Nil(t, mixtape)
	})
}

func TestMixtape_NewPlaylist(t *testing.T) {
	t.Run("has one song", func(t *testing.T) {
		playlist := Playlist{}
		playlist.SongIDs = []string{"1"}

		mixtape := Mixtape{}
		err := mixtape.NewPlaylist(&playlist)
		require.NoError(t, err)
		require.Equal(t, 1, len(mixtape.Playlists))
	})

	t.Run("does not have one song", func(t *testing.T) {
		playlist := Playlist{}

		mixtape := Mixtape{}
		err := mixtape.NewPlaylist(&playlist)
		require.Error(t, err)
		require.Equal(t, ErrNewPlaylistRequiresOneSong, err.Error())
	})

	t.Run("has no playlist yet", func(t *testing.T) {
		playlist := Playlist{}
		playlist.SongIDs = []string{"1"}

		mixtape := Mixtape{}
		err := mixtape.NewPlaylist(&playlist)
		require.NoError(t, err)
		require.Equal(t, 1, len(mixtape.Playlists))
		// we're accessing the first element because we only expect
		// one playlist and we also expect the ID to be 1
		require.Equal(t, "1", mixtape.Playlists[0].ID)
	})

	t.Run("has one playlist", func(t *testing.T) {
		playlist := Playlist{}
		// we need to start from one
		playlist.ID = "1"
		playlist.SongIDs = []string{"1"}

		mixtape := Mixtape{}
		// add one playlist
		mixtape.Playlists = Playlists{&playlist}

		// we'll reuse the same playlist, shouldn't matter since we're
		// generating the ID anyway
		err := mixtape.NewPlaylist(&playlist)
		require.NoError(t, err)
		require.Equal(t, 2, len(mixtape.Playlists))
		// we're accessing the second element because we only expect
		// two playlist and we also expect the ID to be 2
		require.Equal(t, "2", mixtape.Playlists[1].ID)
	})
}

func TestMixtape_RemovePlaylist(t *testing.T) {
	t.Run("playlist exist", func(t *testing.T) {
		// generate the data for testing
		playlistID := "1"
		mixtape := Mixtape{}
		playlist := Playlist{
			ID: playlistID,
		}
		mixtape.Playlists = Playlists{&playlist}

		err := mixtape.RemovePlaylist(playlistID)
		require.NoError(t, err)

		notFound := false
		for i := range mixtape.Playlists {
			playlist := mixtape.Playlists[i]
			if playlist.ID == playlistID {
				notFound = true
				break
			}
		}
		require.False(t, notFound)
	})

	t.Run("playlist does not exist", func(t *testing.T) {
		mixtape := Mixtape{}
		err := mixtape.RemovePlaylist("1")
		require.Error(t, err)
		require.Equal(t, ErrPlaylistNotFound, err.Error())
	})
}

func TestMixtape_AddSongsToPlaylist(t *testing.T) {
	songIDs := []string{"1", "2", "3"}

	t.Run("playlist found", func(t *testing.T) {
		playlistID := "1"
		mixtape := Mixtape{}
		playlist := &Playlist{
			ID: playlistID,
		}
		mixtape.Playlists = Playlists{playlist}

		err := mixtape.AddSongsToPlaylist(playlistID, songIDs)
		require.NoError(t, err)

		playlist = mixtape.Playlists.FindPlaylist(playlistID)
		// check if the songIDs did get added
		require.Equal(t, songIDs, playlist.SongIDs)

	})

	t.Run("playlist not found", func(t *testing.T) {
		mixtape := Mixtape{}
		err := mixtape.AddSongsToPlaylist("1", songIDs)
		require.Error(t, err)
		require.Equal(t, ErrPlaylistNotFound, err.Error())
	})
}

func TestMixtape_ApplyChanges(t *testing.T) {
	t.Run("Changes are valid", func(t *testing.T) {

	})
}
