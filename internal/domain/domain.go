package domain

// Album represents the album that we return to our clients
type Album struct {
	Title      string `json:"title"`
	ArtistName string `json:"artist_name"`
	ImageURL   string `json:"image_url"`
}
