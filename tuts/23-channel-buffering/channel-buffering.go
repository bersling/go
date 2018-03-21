// By default channels are _unbuffered_, meaning that they
// will only accept sends (`chan <-`) if there is a
// corresponding receive (`<- chan`) ready to receive the
// sent value. _Buffered channels_ accept a limited
// number of  values without a corresponding receiver for
// those values.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Here we `make` a channel of strings buffering up to
	// 2 values.
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these
	// values into the channel without a corresponding
	// concurrent receive.
	messages <- "buffered"
	messages <- "channel"
	// messages <- "wuat" // DN: fatal error: all goroutines are asleep - deadlock!

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	// fmt.Println(<-messages)

	// DN Example:
	streamChan := make(chan string, 3)
	streamFn(streamChan)

}

func streamFn(streamChan chan string) {
	fmt.Println("Stream started")
	for i := 0; i < 3; i++ {
		switch i {
		case 0:
			streamChan <- "I"
		case 1:
			streamChan <- "love"
		case 2:
			streamChan <- "you"
		}
		time.Sleep(time.Second)
	}
}
