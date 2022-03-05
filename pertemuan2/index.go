package main

import (
	"fmt"
	"time"
)

func satu() {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Yang cepet")
	}
}

func dua() {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Yang lama")
	}
}

func main () {
	start := time.Now()

	go dua()
	go satu()

	time.Sleep(3000 * time.Millisecond)
	fmt.Println(time.Since(start))
}