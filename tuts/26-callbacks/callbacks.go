// DN: I don't really know whether callbacks are a good idea, but they work...
// DN: This file is entirely written by me.

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hi")
	go readFile(callbackFunction)
	fmt.Println("goodbye (press a key to leave the program)")

	fmt.Scanln()

}

func readFile(cb func(string)) {
	time.Sleep(time.Second)
	cb("Lorem ipsum...")
}

func callbackFunction(returnval string) {
	fmt.Println(returnval)
}
