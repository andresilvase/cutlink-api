package models

type ShortenedURL struct {
	ShortenedURL string `json:"shortened_url"`
}

type FullURL struct {
	URL string `json:"url"`
}
