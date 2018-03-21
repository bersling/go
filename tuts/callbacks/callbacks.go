// DN: I don't really know whether callbacks are a good idea, but they work...

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hi")
	go readFile(func(returnval string) {
		fmt.Println(returnval)
	})
	fmt.Println("goodbye (press a key to leave the program)")

	fmt.Scanln()

}

func readFile(cb func(string)) {
	time.Sleep(time.Second)
	cb("Lorem ipsum...")
}