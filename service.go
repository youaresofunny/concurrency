package main

import "fmt"
import "time"

type IncrementCommand struct {
	CommandModule
}

type CommandModule struct {
	commandModule string
}

func (c CommandModule) CommandModuleName() string {
	return c.commandModule;
}

type Command interface {
	CommandModuleName() string
}

type QueryState struct {
	QueryModule
}

type QueryModule struct {
	queryModule string
}

func (q QueryModule) QueryModuleName() string {
	return q.queryModule
}

type Query interface {
	QueryModuleName() string
}

type Result struct {
	module string
	result int
}

type WorkerChannels struct {
	commandReq <- chan Command
	queryReq <- chan Query
	result chan <- Result
}

type ClientChannels struct {
	command chan <- Command
	query chan <- Query 
	result <- chan Result
	done chan <- bool
}

func module(name string, ctx WorkerChannels) {
	state := 1
	for {
		select {
			case msg := <- ctx.commandReq:
				fmt.Printf("Request %v %T \n", msg, msg)
				switch msg.(type) {
					case IncrementCommand: 
						state += 1
				}
				

			case msg := <- ctx.queryReq:
				fmt.Println("Got query", msg)
				
		}
		result := Result{name, state}
		ctx.result <- result
	}
}

func testClient(ctx ClientChannels) {
	cmd := IncrementCommand{CommandModule{"main"}}
	ctx.command <- cmd
	res := <- ctx.result 
	fmt.Println("Result client", res)

	ctx.command <- cmd
	res = <- ctx.result 
	fmt.Println("Result client", res)

	queryState := QueryState{QueryModule{"name"}}
	ctx.query <- queryState
	res = <- ctx.result 
	fmt.Println("Result client", res)

	time.Sleep(time.Second*10)
	ctx.done <- true
}

func main() {
	command := make(chan Command)
	query := make(chan Query)
	result := make(chan Result)
	done := make(chan bool)

	commandMain := make(chan Command, 5)
	queryMain := make(chan Query, 5)
	resultMain := make(chan Result, 5)

	go module("main", WorkerChannels{commandMain, queryMain, resultMain})
	go testClient(ClientChannels{command, query, result, done})

	for true {
			select {
				case msg := <- command:
					commandMain <- msg
				case msg := <- query:
					queryMain <- msg
				case msg := <- resultMain:
					result <- msg

				case <-done:
					fmt.Println("Done by timeout")
					return;
			}
	}
}