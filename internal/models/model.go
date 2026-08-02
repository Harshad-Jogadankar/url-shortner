package models

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	ShortenedUrl string `json:"shortened_url"`
}

type ResolveRequest struct {
	ShortenedUrl string `json:"shortened_url"`
}

type ResolveResponse struct {
	OriginalUrl string `json:"original_url"`
}
