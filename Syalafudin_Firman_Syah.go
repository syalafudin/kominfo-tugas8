package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Kondisi acak:")

	go printData("Interface1", []interface{}{"bisa1", "bisa2", "bisa3"}, &wg)
	go printData("Interface2", []interface{}{"coba1", "coba2", "coba3"}, &wg)

	wg.Wait()

	fmt.Println("\nKondisi rapih:")

	var mutex sync.Mutex
	var wg2 sync.WaitGroup
	wg2.Add(2)

	go printDataSync("Interface1", []interface{}{"bisa1", "bisa2", "bisa3"}, &wg2, &mutex)
	go printDataSync("Interface2", []interface{}{"coba1", "coba2", "coba3"}, &wg2, &mutex)

	wg2.Wait()
}

func printData(name string, data []interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 4; i++ {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Println(data, i+1)
	}
}

func printDataSync(name string, data []interface{}, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 4; i++ {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		mutex.Lock()
		fmt.Println(data, i+1)
		mutex.Unlock()
	}
}
