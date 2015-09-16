package main

//based on Concurrency patterns (Google IO 2012)

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web1 = fakeSearch("web1")
	Web2 = fakeSearch("web2")
	Web3 = fakeSearch("web3")
	SQL  = fakeSearch("sql")
	SQL2 = fakeSearch("sql2")
	XML  = fakeSearch("xml")
	XML2 = fakeSearch("xml2")
)

type Result string

type Finder func(query string) Result

func fakeSearch(kind string) Finder {
	return func(query string) Result {
		start := time.Now()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		timer := time.Since(start)
		return Result(fmt.Sprintf("%s result for %q : %s\n", kind, query, timer))
	}
}

func Search(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- First(query, Web1, Web2, Web3) }()
	go func() { c <- First(query, SQL, SQL2) }()
	go func() { c <- First(query, XML, XML2) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}

	}
	return
}

func First(query string, replicas ...Finder) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Search("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
