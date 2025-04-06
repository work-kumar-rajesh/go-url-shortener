package memory

import (
	"errors"
	"math/rand"
	"sync"
	"time"

	"github.com/work-kumar-rajesh/go-url-shortner/internal/model"
)

type URLService struct {
	store     map[string]string
	analytics map[string][]model.AnalyticsLog // key is short code
	mutex     sync.RWMutex
}

func NewURLService() *URLService {
	return &URLService{
		store:     make(map[string]string),
		analytics: make(map[string][]model.AnalyticsLog),
	}
}

func (s *URLService) ShortenURL(originalURL string) (string, error) {
	code := generateCode(6)

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store[code] = originalURL

	return code, nil
}

// generateCode creates a random alphanumeric string.
func generateCode(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}
	return string(b)
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

// LogAnalytics records a redirect event.
func (s *URLService) LogAnalytics(code, ip, userAgent string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	logEntry := model.AnalyticsLog{
		ShortCode: code,
		Timestamp: time.Now(),
		IP:        ip,
		UserAgent: userAgent,
	}
	s.analytics[code] = append(s.analytics[code], logEntry)
}

func (s *URLService) GetAnalytics(code string) ([]model.AnalyticsLog, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	logs, exists := s.analytics[code]
	if !exists {
		return nil, errors.New("no analytics found for this short code")
	}
	return logs, nil
}
