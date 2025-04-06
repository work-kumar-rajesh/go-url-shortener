package service

type URLService interface {
	ShortenURL(originalURL string) (string, error)
	ResolveURL(shortCode string) (string, error)
}
