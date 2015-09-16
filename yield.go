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
	c := counter(17)
	c2 := counter(85)

	//Take init val
	fmt.Println(<-c)
	fmt.Println(<-c2)

	//Now there are incremented values in channels
	fmt.Println(<-c)
	fmt.Println(<-c2)
}
