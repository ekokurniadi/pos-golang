package entity

import "time"

type Lokasi struct {
	ID         int
	KodeLokasi string
	Lokasi     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
