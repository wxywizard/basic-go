package main

func Closure(name string) func() string {
	// 闭包
	// name 变量
	// 方法本身
	return func() string {
		return "hello" + name
	}
}

/**
闭包如果使用不当可能会引起内存泄露的问题，
即一个对象被闭包引用的话，它是不会被垃圾回收的
*/
