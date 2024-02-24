package court

import (
	. "l0bby_backend/internal/sportstype"
)

type Court struct {
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	Type    SportsType `json:"type"`
	Address string     `json:"address"`
	Area    string     `json:"area"`
	Phone   string     `json:"phone"`
}
