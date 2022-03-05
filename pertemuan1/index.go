package main 

import (
	"fmt"
	"time"
)

func cetakSesuatu(text string) {
	for i := 0; i < 5; i++ {
		// agar goroutine bisa jalan, klo tidak fungsi cetakSesuatu gk bakal di jalankan
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Dijalankan", text)
	}
}

func main() {
	start := time.Now()
	/*
		======== Belajar Goroutine ========

		Goroutine berfungsi untuk menjalankan
		sebuah function di tempat lain. Jadi
		sebuah function tidak perlu saling menunggu.

		rumus :
		go namaFungsinya()
	*/
	// untuk menjalan sebuah function dengan goroutine
	go cetakSesuatu("Tanpa Goroutine")
	go cetakSesuatu("Dengan Goroutine")

	fmt.Println(time.Since(start))
}