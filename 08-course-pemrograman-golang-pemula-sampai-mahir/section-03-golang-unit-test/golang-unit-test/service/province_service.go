package service

import (
	"errors"
	"golang-unit-test/entity"
	"golang-unit-test/repository"
)

type Province struct {
	repository repository.Province
}

func (p *Province) GetProvince(id int) (*entity.Province, error) {
	result := p.repository.GetProvince(id)
	if result == nil {
		return nil, errors.New("Data not found")
	}
	return result, nil
}
