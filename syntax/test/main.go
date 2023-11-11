package main

func f() any {
	type MyBool bool
	var b MyBool = false
	return b
}

func main() {
	switch f() {
	case false:
		println(1)
	case f():
		println(2)
	case true:
		println(3)
	}
}
