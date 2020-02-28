package mixtape

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Changes represents the structure of the changes that will be made to a
// Mixtape.
type Changes struct {
	Playlists struct {
		New    Playlists `json:"new"`
		Delete struct {
			PlaylistIDs []string `json:"playlist_ids"`
		} `json:"delete"`
		AddSongs []struct {
			PlaylistID string   `json:"playlist_id"`
			SongIDs    []string `json:"song_ids"`
		} `json:"add_songs"`
	} `json:"playlists"`
}

// LoadChanges takes a filePath and opens the file and loads the contents of
// the file to the a reference of a Changes.
func LoadChangesFromFile(filePath string) (*Changes, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	mixtape := Changes{}
	err = json.Unmarshal(b, &mixtape)
	if err != nil {
		return nil, err
	}
	return &mixtape, nil
}
