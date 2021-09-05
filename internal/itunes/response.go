package itunes

type SearchResponse struct {
	ResultCount uint     `json:"resultCount"`
	Results     []Result `json:"results"`
	FromCache   bool
}

type WrapperType string

const (
	WrapperTypeCollection = "collection"
)

type CollectionType string

const (
	CollectionTypeAlbum = "Album"
)

type Result struct {
	WrapperType    WrapperType    `json:"wrapperType"`
	CollectionType CollectionType `json:"collectionType"`
	CollectionName string         `json:"collectionName"`
	ArtistName     string         `json:"artistName"`
	ArtworkURL100  string         `json:"artworkUrl100"`
	ArtworkURL60   string         `json:"artworkUrl60"`
}
