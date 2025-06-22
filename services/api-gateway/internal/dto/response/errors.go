package response

import "time"

type APIErrorResponse struct {
	Title     string    `json:"title" example:"Error title"`
	Details   string    `json:"details" example:"Something went wrong"`
	Timestamp time.Time `json:"timestamp" example:"2025-02-03T15:57:31.17345643+00:00"`
}
