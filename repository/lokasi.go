package repository

import (
	"github.com/ekokurniadi/pos-golang/entity"
	"gorm.io/gorm"
)

type LokasiRepository interface {
	SaveLokasi(lokasi entity.Lokasi) (entity.Lokasi, error)
	UpdateLokasi(lokasi entity.Lokasi) (entity.Lokasi, error)
	FindByIDLokasi(ID int) (entity.Lokasi, error)
	FindAllLokasi() ([]entity.Lokasi, error)
	DeleteByIDLokasi(ID int) (entity.Lokasi, error)
}

type lokasiRepository struct {
	db *gorm.DB
}

func NewLokasiRepository(db *gorm.DB) *lokasiRepository {
	return &lokasiRepository{db}
}

func (r *lokasiRepository) SaveLokasi(lokasi entity.Lokasi) (entity.Lokasi, error) {
	err := r.db.Create(&lokasi).Error
	if err != nil {
		return lokasi, err
	}
	return lokasi, nil

}
func (r *lokasiRepository) FindByIDLokasi(ID int) (entity.Lokasi, error) {
	var lokasi entity.Lokasi
	err := r.db.Where("id = ? ", ID).Find(&lokasi).Error
	if err != nil {
		return lokasi, err
	}
	return lokasi, nil

}
func (r *lokasiRepository) UpdateLokasi(lokasi entity.Lokasi) (entity.Lokasi, error) {
	err := r.db.Save(&lokasi).Error
	if err != nil {
		return lokasi, err
	}
	return lokasi, nil

}
func (r *lokasiRepository) FindAllLokasi() ([]entity.Lokasi, error) {
	var lokasis []entity.Lokasi
	err := r.db.Find(&lokasis).Error
	if err != nil {
		return lokasis, err
	}
	return lokasis, nil

}
func (r *lokasiRepository) DeleteByIDLokasi(ID int) (entity.Lokasi, error) {
	var lokasi entity.Lokasi
	err := r.db.Where("id = ? ", ID).Delete(&lokasi).Error
	if err != nil {
		return lokasi, err
	}
	return lokasi, nil

}

//Generated by Micagen at 08 Desember 2021
