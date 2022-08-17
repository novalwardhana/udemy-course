package belajar_golang_context

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestMembuatContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestMakeMultipleContext(t *testing.T) {
	contextBackgroundArr := []context.Context{}
	contextTodoArr := []context.Context{}
	iteration := 10

	for i := 0; i < iteration; i++ {
		bg := context.Background()
		todo := context.TODO()

		contextBackgroundArr = append(contextBackgroundArr, bg)
		contextTodoArr = append(contextTodoArr, todo)
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, data := range contextBackgroundArr {
			fmt.Println("Context background: ", data)
		}
	}()
	go func() {
		defer wg.Done()
		for _, data := range contextTodoArr {
			fmt.Println("Context todo: ", data)
		}
	}()
	wg.Wait()

	fmt.Println("Finish")
}
