package main

import (
	// "time"
	"fmt"
)
var myChannel = make(chan string)
func main() {
	// myChannel := make(chan string)

	go isiChannel("Test")

	for channel := range myChannel {
		fmt.Println(channel)
	}
}

func isiChannel(data string) {
	for i := 0; i < 5; i++ {
		myChannel <- data
	}

	// unutk menuutu sebuah channel
	// lakukan ini juga ada error : all goroutines are asleep - deadlock
	// dimana channel akan terus di cetak padahal channelnyya udh habis
	// maksudnya ada 3 channel, tapi malah di looping utk di tampilakn 5 kali
	close(myChannel)
}
