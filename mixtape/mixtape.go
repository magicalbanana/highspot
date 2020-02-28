package mixtape

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

const (
	ErrNewPlaylistRequiresOneSong = "New playlist requires at least one song"
	ErrPlaylistNotFound           = "Playlist not found"
)

// Mixtape represents the primary data structure of a "mixtape" record
type Mixtape struct {
	Users     Users     `json:"users"`
	Playlists Playlists `json:"playlists"`
	Songs     Songs     `json:"songs"`
}

// LoadMixtape takes a filePath and opens the file and loads the contents of
// the file to the a reference of a Mixtape.
func LoadMixtapeFromFile(filePath string) (*Mixtape, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	mixtape := Mixtape{}
	err = json.Unmarshal(b, &mixtape)
	if err != nil {
		return nil, err
	}
	return &mixtape, nil
}

// NewPlaylist adds the givne playlist to the Playlists. The playlist must
// have at least one song in order to be added. If no song is given, it returns
// an error. It does not check whether the song exists or not.
func (m *Mixtape) NewPlaylist(playlist *Playlist) error {
	if len(playlist.SongIDs) == 0 {
		return errors.New(ErrNewPlaylistRequiresOneSong)
	}

	if m.Playlists == nil || len(m.Playlists) < 1 {
		m.Playlists = make(Playlists, 0)
	}

	// generate ID for Playlist
	id, err := m.Playlists.generateNextID()
	if err != nil {
		return err
	}
	playlist.ID = id
	m.Playlists = append(m.Playlists, playlist)

	return nil
}

// RemovePlaylist removes a playlist with the given playlistID. If the
// playlistID is not matched it returns an error.
func (m *Mixtape) RemovePlaylist(playlistID string) error {
	// The reason this is a pointer is so that we can tell if the
	// playlist is found because in "go", the zero value will be 0 and we
	// as we know, arrays start at zero. ;)
	var removeFromIndex *int

	for i := range m.Playlists {
		playlist := m.Playlists[i]
		if playlist.ID == playlistID {
			removeFromIndex = &i
			break
		}
	}

	if removeFromIndex == nil {
		return errors.New(ErrPlaylistNotFound)
	}

	m.Playlists = append(m.Playlists[:*removeFromIndex], m.Playlists[*removeFromIndex+1:]...)
	return nil
}

// AddSongsToPlaylist adds the given songIDs to the playlist with the
// playlistID. This function does not take into account duplicate songIDs in
// the same playlist.
func (m *Mixtape) AddSongsToPlaylist(playlistID string, songIDs []string) error {
	playlist := m.Playlists.FindPlaylist(playlistID)
	if playlist == nil {
		return errors.New(ErrPlaylistNotFound)
	}

	for i := range songIDs {
		playlist.AddSong(songIDs[i])
	}

	return nil
}

// ApplyChanges takes the given changes and applies them to the current
// Mixtape
func (m *Mixtape) ApplyChanges(changes *Changes) error {
	// add new playlist
	for i := range changes.Playlists.New {
		playlist := changes.Playlists.New[i]
		err := m.NewPlaylist(playlist)
		if err != nil {
			return err
		}
	}

	// remove playlist
	for i := range changes.Playlists.Delete.PlaylistIDs {
		err := m.RemovePlaylist(changes.Playlists.Delete.PlaylistIDs[i])
		if err != nil {
			return err
		}
	}

	// add song to playlist
	for i := range changes.Playlists.AddSongs {
		newSong := changes.Playlists.AddSongs[i]
		err := m.AddSongsToPlaylist(newSong.PlaylistID, newSong.SongIDs)
		if err != nil {
			return err
		}
	}

	return nil
}
