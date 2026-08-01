package models

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	ShortenedUrl string `json:"shortened_url"`
}
