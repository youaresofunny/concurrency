package main

//based on Concurrency patterns (Google IO 2012)

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan string)
	j := boring("Joe", quit)
	b := boring("Bob", quit)
	for i := rand.Intn(20); i >= 0; i-- {
		fmt.Println(<-j)
		fmt.Println(<-b)
	}
	quit <- "Bye!"
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func boring(msg string, quit chan string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
			case <-quit:
				cleanup()
				quit <- "See you!"
				return
			}
		}
	}()
	return c // Return the channel to the caller.
}

func cleanup() {}
