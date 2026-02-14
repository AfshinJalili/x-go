// Package selectpkg demonstrates the select statement for channel operations.
// This file shows selecting between multiple channels.
package selectpkg

import (
	"fmt"
	"time"
)

// SelectChannel demonstrates selecting between two channels.
// The select statement will choose the first channel that becomes ready.
func SelectChannel() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}
