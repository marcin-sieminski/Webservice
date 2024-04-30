package data

import (
	"time"
)

type Item struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Version   int32     `json:"-"`
}
