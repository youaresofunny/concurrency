package main

//Simple int generator (yield using chan)

import (
	"fmt"
)

func counter(n int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			c <- n
			n++
		}
	}()
	return c
}

func main() {
	c1 := counter(4)
	c2 := counter(44)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
}
