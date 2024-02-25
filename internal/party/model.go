package party

import (
	. "l0bby_backend/internal/court"
	"time"
)

type Party struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Court     Court     `json:"court"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Members   []string  `json:"members"`
}
