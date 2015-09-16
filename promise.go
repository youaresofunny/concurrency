package main

import (
	"fmt"
	"sync"
)

type PromiseDelivery chan interface{}

type Promise struct {
	sync.RWMutex
	value   interface{}
	waiters []PromiseDelivery
}

func (p *Promise) Deliver(value interface{}) {
	p.Lock()
	defer p.Unlock()
	p.value = value
	for _, w := range p.waiters {
		locW := w
		go func() {
			locW <- value
		}()
	}
}

func (p *Promise) Value() interface{} {
	if p.value != nil {
		return p.value
	}

	delivery := make(PromiseDelivery)
	p.waiters = append(p.waiters, delivery)
	return <-delivery
}

func NewPromise() *Promise {
	return &Promise{
		value:   nil,
		waiters: []PromiseDelivery{},
	}
}

func main() {
	v := NewPromise()
	go func() {
		v.Deliver(42)
	}()
	go func() {
		fmt.Println(v.Value().(int))
	}()
	fmt.Println(v.Value().(int))
}
