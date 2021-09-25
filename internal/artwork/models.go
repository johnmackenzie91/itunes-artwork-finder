package artwork

// SearchResponse is the response that will be returned by all finders. This bring the response into our domain
type SearchResponse struct {
	ResultCount uint `json:"resultCount"`
	Results     []Result
}

// Result represents each single result
type Result struct {
	Artist         string
	CollectionName string
	Link           string
}
