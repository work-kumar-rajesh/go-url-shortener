package memory

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

// URLService is an in-memory implementation of the URLService interface.
type URLService struct {
	store map[string]string
	mutex sync.RWMutex
}

func NewURLService() *URLService {
	return &URLService{
		store: make(map[string]string),
	}
}

func (s *URLService) ShortenURL(originalURL string) (string, error) {
	code := generateCode(6)

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store[code] = originalURL

	return code, nil
}

func (s *URLService) ResolveURL(code string) (string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	url, exists := s.store[code]
	if !exists {
		return "", errors.New("not found")
	}
	return url, nil
}

// generateCode creates a random alphanumeric string of a given length.
func generateCode(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}
	return string(b)
}
