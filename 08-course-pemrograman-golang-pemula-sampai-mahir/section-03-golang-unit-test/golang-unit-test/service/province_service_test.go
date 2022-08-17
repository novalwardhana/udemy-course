package service

import (
	"golang-unit-test/entity"
	"golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var provinceRepository = repository.ProvinceMock{Mock: mock.Mock{}}
var provinceService = Province{repository: &provinceRepository}

func TestProvince(t *testing.T) {
	provinceRepository.Mock.On("GetProvince", 12).Return(entity.Province{
		ID:        12,
		CountryID: 1,
		Code:      "P0012",
		Name:      "Yogyakarta",
	})
	result, err := provinceService.GetProvince(12)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestProvinceNil(t *testing.T) {
	provinceRepository.Mock.On("GetProvince", 13).Return(nil)
	result, err := provinceService.GetProvince(13)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}
