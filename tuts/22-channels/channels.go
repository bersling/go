// DN: Sounds simple
// DN: As for everything, there are some critics:
// DN: https://www.jtolio.com/2016/03/go-channels-are-bad-and-you-should-feel-bad/
// DN: I can't judge yet whether this is justified or not

// _Channels_ are the pipes that connect concurrent
// goroutines. You can send values into channels from one
// goroutine and receive those values into another
// goroutine.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Create a new channel with `make(chan val-type)`.
	// Channels are typed by the values they convey.
	messages := make(chan string)

	// _Send_ a value into a channel using the `channel <-`
	// syntax. Here we send `"ping"`  to the `messages`
	// channel we made above, from a new goroutine.
	go func() { messages <- "ping" }()

	// The `<-channel` syntax _receives_ a value from the
	// channel. Here we'll receive the `"ping"` message
	// we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)

	// DN Example:
	countdownChannel := make(chan int)

	go countdownFn(countdownChannel)

	// The go routine will stop at the line with the channel piping, because the channel isn't ready yet
	time.Sleep(time.Millisecond * 2000)

	i := <-countdownChannel
	fmt.Println(i)
	j := <-countdownChannel
	fmt.Println(j)
	k := <-countdownChannel
	fmt.Println(k)
	fmt.Println("Hm, I need to wait for you guys?")

	// DN Example 2:
	countdownChannel2 := make(chan int)

	go countdownFn(countdownChannel2)

	go func() {
		i := <-countdownChannel2 // DN: Ok, that's not so nice: i is redefined, but GO doesn't warn
		fmt.Println(i)
		j := <-countdownChannel2
		fmt.Println(j)
		k := <-countdownChannel2
		fmt.Println(k)
	}()
	fmt.Println("Ok you guys can do your stuff, but I'm in a hurry")

	fmt.Scanln()

}

func countdownFn(countdownChannel chan int) {
	fmt.Println("Started Countdown!")
	for i := 10; i >= 0; i-- {
		countdownChannel <- i // ... only doing this once channel is ready
		time.Sleep(time.Second)
	}
}
