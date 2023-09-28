package main

import (
	"fmt"
	"time"
)

func main() {
	println("Hello, Go!")
	fmt.Println(DeferReturnV0())
	fmt.Println(DeferReturnV1())
}

func Test3() (a int) {
	defer func() {
		a = 0
	}()
	return a
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

func DeferReturnV0() int {
	a := 1
	defer func() {
		a = 0
	}()
	return a
}

func DeferReturnV1() (a int) {
	a = 1
	defer func() {
		a = 0
	}()
	return a
}
