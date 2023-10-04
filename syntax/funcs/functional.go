package main

func MyFunc3() {
	print("hello MyFunc3")
}

func Func4() {
	fn := Func1
	s, err := fn(1, 2, 3)
	println(s, err)
}

func Func5() {
	fn := func(name string) string {
		return "hello" + name
	}
	s := fn("wer")
	println(s)
}

func Func6() (func(name string) string, func(se int) int) {
	fn1 := func(name string) string {
		return "hello" + name
	}
	fn2 := func(se int) int {
		return se * se
	}
	return fn1, fn2
}

func Func6Invoke() {
	fn1, fn2 := Func6()
	println(fn1("wer"))
	println(fn2(2))
}

func SwitchFunc(i int) {
	switch i {
	case 23:
		println("23")
		fallthrough
	case 24:
		println("24")

	}
}
