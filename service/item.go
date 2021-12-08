package service

import (
	"github.com/ekokurniadi/pos-golang/entity"
	"github.com/ekokurniadi/pos-golang/input"
	"github.com/ekokurniadi/pos-golang/repository"
)

type ItemService interface {
	ItemServiceGetAll() ([]entity.Item, error)
	ItemServiceGetByID(inputID input.InputIDItem) (entity.Item, error)
	ItemServiceCreate(input input.ItemInput) (entity.Item, error)
	ItemServiceUpdate(inputID input.InputIDItem, inputData input.ItemInput) (entity.Item, error)
	ItemServiceDeleteByID(inputID input.InputIDItem) (bool, error)
	ItemServiceSearch(input input.InputFilter) ([]entity.Item, error)
}
type itemService struct {
	repository repository.ItemRepository
}

func NewItemService(repository repository.ItemRepository) *itemService {
	return &itemService{repository}
}
func (s *itemService) ItemServiceCreate(input input.ItemInput) (entity.Item, error) {
	item := entity.Item{}
	item.KodeProduk = input.KodeProduk
	item.Jenis = input.Jenis
	item.Kategori = input.Kategori
	item.Satuan = input.Satuan
	item.NamaProduk = input.NamaProduk
	item.Keterangan = input.Keterangan
	item.Lokasi = input.Lokasi
	item.Gambar = input.Gambar
	item.HargaJual = input.HargaJual
	item.Stok = input.Stok
	newItem, err := s.repository.SaveItem(item)
	if err != nil {
		return newItem, err
	}
	return newItem, nil
}
func (s *itemService) ItemServiceUpdate(inputID input.InputIDItem, inputData input.ItemInput) (entity.Item, error) {
	item, err := s.repository.FindByIDItem(inputID.ID)
	if err != nil {
		return item, err
	}
	item.KodeProduk = inputData.KodeProduk
	item.Jenis = inputData.Jenis
	item.Kategori = inputData.Kategori
	item.Satuan = inputData.Satuan
	item.NamaProduk = inputData.NamaProduk
	item.Keterangan = inputData.Keterangan
	item.Lokasi = inputData.Lokasi
	item.Gambar = inputData.Gambar
	item.HargaJual = inputData.HargaJual
	item.Stok = inputData.Stok

	updatedItem, err := s.repository.UpdateItem(item)

	if err != nil {
		return updatedItem, err
	}
	return updatedItem, nil
}
func (s *itemService) ItemServiceGetByID(inputID input.InputIDItem) (entity.Item, error) {
	item, err := s.repository.FindByIDItem(inputID.ID)
	if err != nil {
		return item, err
	}
	return item, nil
}
func (s *itemService) ItemServiceSearch(input input.InputFilter) ([]entity.Item, error) {
	items, err := s.repository.SearchItem(input)
	if err != nil {
		return items, err
	}
	return items, nil
}

func (s *itemService) ItemServiceGetAll() ([]entity.Item, error) {
	items, err := s.repository.FindAllItem()
	if err != nil {
		return items, err
	}
	return items, nil
}
func (s *itemService) ItemServiceDeleteByID(inputID input.InputIDItem) (bool, error) {
	_, err := s.repository.FindByIDItem(inputID.ID)
	if err != nil {
		return false, err
	}
	_, err = s.repository.DeleteByIDItem(inputID.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}

//Generated by Micagen at 08 Desember 2021
