package entity

import "time"

type Item struct {
	ID         int
	KodeProduk string
	Jenis      string
	Kategori   string
	Satuan     string
	NamaProduk string
	Keterangan string
	Lokasi     string
	Gambar     string
	HargaJual  int
	Stok       int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
