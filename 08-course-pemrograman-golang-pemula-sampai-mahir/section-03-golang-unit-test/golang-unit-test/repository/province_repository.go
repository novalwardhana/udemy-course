package repository

import (
	"golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type ProvinceMock struct {
	Mock mock.Mock
}

type Province interface {
	GetProvince(id int) *entity.Province
}

func (p *ProvinceMock) GetProvince(id int) *entity.Province {
	result := p.Mock.Called(id)
	if result.Get(0) == nil {
		return nil
	}
	province := result.Get(0).(entity.Province)
	return &province
}
