package entity

import "time"

type Kategori struct {
	ID        int
	Kategori  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
