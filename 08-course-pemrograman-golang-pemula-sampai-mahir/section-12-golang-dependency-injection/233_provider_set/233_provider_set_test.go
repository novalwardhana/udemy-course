package simple

import (
	"fmt"
	"testing"
)

func TestProviderRepository(t *testing.T) {
	repository := InitializedRepository()
	fmt.Println(repository)
}

func TestProviderService(t *testing.T) {
	service := InitializedService(true)
	fmt.Println(service)
}
