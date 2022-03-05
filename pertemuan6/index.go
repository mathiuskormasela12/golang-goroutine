package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// memberitahu ada sebuah goroutine yg harus di tunggu
	wg.Add(1)
	go jalan(&wg)
	// untuk menunggu hingga goroutine berhasil dijalankan
	wg.Wait()
}

func jalan(wg *sync.WaitGroup) {
	fmt.Println("hello")
	// untuk menyelesaikan pemanggilan goroutine
	wg.Done()
}
