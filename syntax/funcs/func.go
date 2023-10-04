package main

func main() {
	//Invoke()
	//DeferClosure()
	//DeferClosureV1()
	DeferClosureLoopV3()
}

func Invoke() {
	_, _ = Func0()
	str1, err1 := Func1(1, 2, 3)
	println(str1, err1)
}

func Func0() (string, error) {
	return "", nil
}

func Func1(arr ...int) (string, error) {
	arr[0] = 1
	return "", nil
}

/**
 * 栈 go routine 私有的 堆 go routine 共享的
 */
