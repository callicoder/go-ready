package model

import (
	"time"
)

type Audit struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
