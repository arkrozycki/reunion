package message

import "time"

// Channel struct
type Channel struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	UserIDs   []int64   `json:"user_i_ds"`
	CreatedAt time.Time `json:"created_at"`
}
