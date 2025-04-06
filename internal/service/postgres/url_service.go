package postgres

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/work-kumar-rajesh/go-url-shortner/internal/model"
)

const shortCodeLength = 6
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type PostgresService struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *PostgresService {
	return &PostgresService{db: db}
}

func generateCode(length int) string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}
	return string(b)
}

func (s *PostgresService) ShortenURL(originalURL string) (string, error) {
	for i := 0; i < 3; i++ {
		code := generateCode(shortCodeLength)

		_, err := s.db.Exec(
			"INSERT INTO urls (short_code, original_url) VALUES ($1, $2)",
			code, originalURL,
		)

		if err == nil {
			return code, nil
		}
	}

	return "", errors.New("failed to generate unique short code")
}

func (s *PostgresService) ResolveURL(code string) (string, error) {
	var originalURL string
	err := s.db.Get(
		&originalURL,
		"SELECT original_url FROM urls WHERE short_code = $1",
		code,
	)

	if err != nil {
		return "", err
	}

	return originalURL, nil
}

func (s *PostgresService) LogAnalytics(shortCode, ipAddress, userAgent string) {
	_, err := s.db.Exec(
		"INSERT INTO analytics (short_code, ip_address, user_agent) VALUES ($1, $2, $3)",
		shortCode, ipAddress, userAgent,
	)
	if err != nil {
		log.Printf("Failed to insert analytics: %v", err)
	}
}

func (s *PostgresService) GetAnalytics(code string) ([]model.AnalyticsLog, error) {
	var logs []model.AnalyticsLog

	err := s.db.Select(&logs,
		"SELECT * FROM analytics WHERE short_code = $1 ORDER BY created_at DESC",
		code,
	)

	if err != nil {
		return nil, err
	}

	return logs, nil
}
