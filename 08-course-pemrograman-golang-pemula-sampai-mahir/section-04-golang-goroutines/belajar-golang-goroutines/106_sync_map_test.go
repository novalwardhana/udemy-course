package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func AddSyncMapInt(dataMap *sync.Map, value int, wg *sync.WaitGroup) {
	defer wg.Wait()

	dataMap.Store(value, value)
}

func TestSyncMapInt(t *testing.T) {
	dataMap := &sync.Map{}
	wg := &sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go AddSyncMapInt(dataMap, i, wg)
	}
	fmt.Println("Success generate map data")

	dataMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}

func SyncMapStore(dataMap *sync.Map, param int, wg *sync.WaitGroup) {
	defer wg.Done()

	key := fmt.Sprintf("index-%d", param)
	value := float64(param) / float64(param+1)
	dataMap.Store(key, value)
}

func SyncMapLoad(dataMap *sync.Map, key string, wg *sync.WaitGroup) {
	defer wg.Done()

	value, ok := dataMap.Load(key)
	if ok {
		fmt.Println(key, ":", value)
	} else {
		fmt.Println(key, ":", value)
	}
}

func SyncMapDelete(dataMap *sync.Map, key string, wg *sync.WaitGroup) {
	defer wg.Done()

	dataMap.Delete(key)
}

func TestSyncMapAddDelete(t *testing.T) {
	var dataMap = &sync.Map{}
	var wg = &sync.WaitGroup{}
	iteration := 100

	fmt.Println("Start store to map")
	for i := 1; i <= iteration; i++ {
		wg.Add(1)
		go SyncMapStore(dataMap, i, wg)
	}
	wg.Wait()
	fmt.Println("Finish store to map")

	fmt.Println("Start load map")
	for i := 1; i <= iteration; i++ {
		wg.Add(1)
		key := fmt.Sprintf("index-%d", i)
		go SyncMapLoad(dataMap, key, wg)
	}
	wg.Wait()
	fmt.Println("Finish load map")

	fmt.Println("Start delete map")
	for i := 1; i <= iteration; i++ {
		wg.Add(1)
		key := fmt.Sprintf("index-%d", i)
		go SyncMapDelete(dataMap, key, wg)
	}
	wg.Wait()
	fmt.Println("Finish delete map")

	fmt.Println("Start load map")
	for i := 1; i <= iteration; i++ {
		wg.Add(1)
		key := fmt.Sprintf("index-%d", i)
		go SyncMapLoad(dataMap, key, wg)
	}
	wg.Wait()
	fmt.Println("Finish load map")
}
