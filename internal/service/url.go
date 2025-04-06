package service

import "github.com/work-kumar-rajesh/go-url-shortner/internal/model"

type URLService interface {
	ShortenURL(originalURL string) (string, error)
	ResolveURL(shortCode string) (string, error)
	LogAnalytics(shortCode, ip, userAgent string)
	GetAnalytics(shortCode string) ([]model.AnalyticsLog, error)
}
