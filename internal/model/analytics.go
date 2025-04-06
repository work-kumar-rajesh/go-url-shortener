package model

import "time"

// AnalyticsLog represents a single redirect event.
type AnalyticsLog struct {
	ShortCode string    `json:"short_code"`
	Timestamp time.Time `json:"timestamp"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"user_agent"`
}
