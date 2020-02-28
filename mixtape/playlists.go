package mixtape

import (
	"sort"
	"strconv"
)

// Playlist ...
type Playlist struct {
	ID      string   `json:"id"`
	UserID  string   `json:"user_id"`
	SongIDs []string `json:"song_ids"`
}

// AddSong adds the given songID to the SongsID field. This does not check if
// a songID exist
func (p *Playlist) AddSong(songID string) {

	if p.SongIDs == nil || len(p.SongIDs) == 0 {
		p.SongIDs = make([]string, 0)
	}
	p.SongIDs = append(p.SongIDs, songID)
}

// Playists is a collection of Playlist references. The reason this is a
// reference is because when we add songs to a playlist on Mixtape we want to
// be able to modify the songIDs of the a Playlist by reference.
type Playlists []*Playlist

func (p Playlists) FindPlaylist(playlistID string) *Playlist {
	for i := range p {
		playlist := p[i]
		if playlist.ID == playlistID {
			return playlist
		}
	}

	return nil
}

// generateNextID generates the next ID from the playlist collection. We use
// this when we're adding a new playlist to the playlists collection.
func (p Playlists) generateNextID() (string, error) {
	if len(p) < 1 {
		return "1", nil
	}

	idsInt := make([]int, 0)
	for i := range p {
		idInt, err := strconv.Atoi(p[i].ID)
		if err != nil {
			return "", err
		}
		idsInt = append(idsInt, idInt)
	}

	sort.Ints(idsInt)

	return strconv.Itoa((idsInt[len(idsInt)-1] + 1)), nil
}
