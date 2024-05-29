package main

import "fmt"

type core struct {
	name  string
	value string
}

type OptionFunc func(*core)

func NewCore(options ...OptionFunc) *core {
	c := &core{
		name:  "default",
		value: "default",
	}

	for _, option := range options {
		option(c)
	}
	return c
}

func WithName(name string) OptionFunc {
	return func(c *core) {
		c.name = name
	}
}

func WithValue(value string) OptionFunc {
	return func(c *core) {
		c.value = value
	}
}

func main() {
	c := NewCore(
		WithName("123"), WithValue("890"),
	)
	fmt.Println(c)
}
