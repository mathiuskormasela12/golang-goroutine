package main

import (
	"fmt"
	"time"
)

// membuat channel
var newChannel = make(chan string)

func main() {
	var data = [3]string{"Javascript", "Go", "Python"}

	go jalan(data)
	go lari()

	time.Sleep(500 * time.Millisecond)
}

func jalan(data [3]string) {
	for _, singleData := range data {
		if singleData == "Go" {
			time.Sleep(500 * time.Millisecond)
			// mengassign nilai singleData ke channel newChannel
			newChannel <- singleData
		}
	}
}

func lari() {
	// untuk menampilkan channel harus menggunakan for range loop
	for channel := range newChannel {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(channel)
	}
}