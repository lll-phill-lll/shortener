package api

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	HashedURL string `json:"hashed_url"`
}