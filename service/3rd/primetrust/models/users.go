package models

const (
	UserType = "users"
)

type UserAttributes struct {
	Disabled        bool   `json:"disabled"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	CurrentPassword string `json:"current-password"`
}

type UserData struct {
	ID            string         `json:"id,omitempty"`
	Type          string         `json:"type"`
	Attributes    UserAttributes `json:"attributes"`
	Links         Links          `json:"links"`
	Relationships Relationships  `json:"relationships"`
}

type User struct {
	Data UserData `json:"data"`
}

type UsersResponse struct {
	CollectionResponse
	Data []UserData `json:"data"`
}
