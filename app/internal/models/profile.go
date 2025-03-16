package models

import "time"

type Profile struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name"`
	Longitude float32   `json:"longitude"`
	Latitude  float32   `json:"latitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
