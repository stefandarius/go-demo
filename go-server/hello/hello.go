package hello

import (
	"fmt"
)

type Hello struct{}

func New() Hello {
	return Hello{}
}

func (h Hello) SayHello() (string, error) {
	fmt.Println("Hello World!")
	return "Hello World!", nil
}

func (h Hello) SayHelloTo(name string) (string, error) {
	fmt.Println("Hello " + name + "!")
	return "Hello " + name + "!", nil
}
