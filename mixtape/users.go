package mixtape

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Users []User
