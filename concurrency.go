// This file contain channel & goroutine example
// channelName := make(chan dataType,length)
// ci := make(chan int)
// cs := make(chan string)
// ch <- v    send v to channel ch.
// v := <-ch  receive data from ch, and assign to v

package main

import (
	"fmt"
	"time"
)

// Print txt 5 times
func sayRepeat(txt string) {
	for i := 0; i < 10; i++ {
		// it sleeps for 1 seconds everytime start a goroutine and return back to this position start
		time.Sleep(100 * time.Microsecond)
		fmt.Println(txt)
	}
}

// func vulnService(interval int, channel chan int) {}
func main() {
	go sayRepeat("ping victim") // create a new goroutine
	sayRepeat("reply from victim")

	// Create a new channel of type string
	payload := make(chan string)

	// Send a string value to message channel
	go func() { payload <- "brute-force -f user.txt -h 1.2.3.4 -i 10" }()

	// Receive string to a variable
	victim := <-payload
	fmt.Println(victim)

	payload2 := make(chan string, 2) // Will allow just for 2 length
	payload2 <- "petya.py --enc /boot --host 1.2.3.4 --rhost 10.132.12.1 -k xcert.pem"
	payload2 <- "echo 'Your /boot sector in your system has been encrypted, pay me some btc so i can buy a cup of coffe'"

	payload2 <- "hey!!"

	victim2 := <-payload2
	fmt.Println(victim2)
	fmt.Println(victim2)
	fmt.Println(victim2) // Will trigger error because we just send the third length
}
