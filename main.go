package main

import (
	"fmt"
	"strings"
	"sync"
)

type Show interface {
	Print() []string
}

type Data struct {
	DataSlice []string
}

func (d Data) Print(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Printf("%s %d\n", d.DataSlice, i)
}

func (d Data) PrinLock(wg *sync.WaitGroup, i int, mtx *sync.Mutex) {
	defer mtx.Unlock()
	defer wg.Done()
	fmt.Printf("%s %d\n", d.DataSlice, i)
}

func main() {
	coba := Data{
		DataSlice: []string{"coba1", "coba2", "coba3"},
	}

	bisa := Data{
		DataSlice: []string{"bisa1", "bisa2", "bisa3"},
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(1)
	go withoutMutex(&coba, &bisa, &wg)
	wg.Wait()

	fmt.Println(strings.Repeat("=", 50))

	wg.Add(1)
	go withMutex(&coba, &bisa, &wg, &mutex)
	wg.Wait()
}

func withoutMutex(coba *Data, bisa *Data, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("without mutex:")
	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go coba.Print(wg, i)
		go bisa.Print(wg, i)
	}

}

func withMutex(coba, bisa *Data, wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()
	fmt.Println("with mutex:")
	for i := 1; i <= 4; i++ {
		wg.Add(2)
		mtx.Lock()
		go coba.PrinLock(wg, i, mtx)
		mtx.Lock()
		go bisa.PrinLock(wg, i, mtx)
	}
}
