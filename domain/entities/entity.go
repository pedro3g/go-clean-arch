package entities

import (
	"time"
)

type Entity struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
