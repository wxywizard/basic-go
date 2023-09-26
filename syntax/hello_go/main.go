package main

import (
	"fmt"
	"time"
)

func main() {
	println("Hello, Go!")
	fmt.Println(Test3(4))
}

func Hello() {
	list := []*Demo{{"a"}, {"b"}}
	for _, item := range list {
		go func(item *Demo) {
			fmt.Println("name" + item.Name)
		}(item)
	}
	time.Sleep(100 * time.Millisecond)
}

type Demo struct {
	Name string
}

func Test1() {
	a := 1
	defer fmt.Println(a)
}

func Test2() int {
	a := 1
	defer func() {
		a = 0
	}()
	return a
}

func Test3(a int) int {
	defer func(a int) {
		a = 0
	}(a)
	return a
}
