package entity

import "time"

type Jenis struct {
	ID        int
	Jenis     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
