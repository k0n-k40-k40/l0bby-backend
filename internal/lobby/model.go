package lobby

import (
	. "l0bby_backend/internal/sportstype"
)

type Lobby struct {
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	Type    SportsType `json:"type"`
	Members []string   `json:"members"`
}
