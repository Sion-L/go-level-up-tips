package main

import "fmt"

type Handler interface {
	Handle(request string)
	SetNext(handler Handler)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

type ConcreteOne struct {
	BaseHandler
}

func (c *ConcreteOne) Handle(request string) {
	if request == "one" {
		fmt.Printf("handler request %s\n", request)
	} else if c.next != nil {
		c.next.Handle(request)
	}
}

type ConcreteTwo struct {
	BaseHandler
}

func (c *ConcreteTwo) Handle(request string) {
	if request == "two" {
		fmt.Printf("handler request %s\n", request)
	} else if c.next != nil {
		c.next.Handle(request)
	}
}

type ConcreteThree struct {
	BaseHandler
}

func (c *ConcreteThree) Handle(request string) {
	if request == "three" {
		fmt.Printf("handler request %s\n", request)
	} else if c.next != nil {
		c.next.Handle(request)
	}
}

type ConcreteOther struct {
	BaseHandler
}

func (c *ConcreteOther) Handle(request string) {
	if request != "" {
		fmt.Printf("other: %s\n", request)
	} else if c.next != nil {
		c.next.Handle(request)
	}
}

func main() {

	c1 := &ConcreteOne{}
	c2 := &ConcreteTwo{}
	c3 := &ConcreteThree{}
	c4 := &ConcreteOther{}

	// 构建责任链
	c1.SetNext(c2)
	c2.SetNext(c3)
	c3.SetNext(c4)
	c4.SetNext(c1)

	for _, v := range []string{"one", "two", "three", "four", "ccc", "ddd"} {
		c1.Handle(v)
	}
}
