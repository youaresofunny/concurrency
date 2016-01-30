package main

import (
	"fmt"
)

type Observable struct {
	Value int
}

type MultiplierCallback struct {
	Value int
}

func (m *MultiplierCallback) Run(o *Observable) {
	o.Value *= m.Value
}

type DivisionCallback struct {
	Value int
}

func (d *DivisionCallback) Run(o *Observable) {
	o.Value = o.Value / d.Value
}

type Callback interface {
	Run(h *Observable)
}

type Observer struct {
	callbacks []Callback
}

func (o *Observer) Add(c Callback) {
	o.callbacks = append(o.callbacks, c)
}

func (o *Observer) Process(oe *Observable) {
	for _, c := range o.callbacks {
		c.Run(oe)
	}
}

func main() {
	oe := Observable{1}
	o := Observer{}
	o.Add(&MultiplierCallback{500})
	o.Add(&DivisionCallback{2})
	o.Process(&oe)
	fmt.Println(oe.Value)
}
