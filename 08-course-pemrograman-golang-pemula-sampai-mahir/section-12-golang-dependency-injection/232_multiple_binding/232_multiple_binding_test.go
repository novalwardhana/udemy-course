package simple

import (
	"fmt"
	"testing"
)

func TestUserService(t *testing.T) {
	userService := InitializedUserService(true)
	fmt.Println(userService)
}
