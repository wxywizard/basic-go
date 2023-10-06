package main

import "fmt"

func main() {

	sliceV1 := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", sliceV1)

	// 删除索引为2的元素
	sliceV1 = deleteAtIndexV1(sliceV1, 2)
	fmt.Println("删除后的切片:", sliceV1)

	sliceV2 := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", sliceV2)

	// 删除索引为2的元素
	sliceV2 = deleteAtIndexV2(sliceV2, 2)
	fmt.Println("删除后的切片:", sliceV2)

	sliceV3 := []int{1, 2, 3, 4, 5}
	fmt.Println("原始整数切片:", sliceV3)
	sliceV3 = deleteAtIndexV3(sliceV3, 2)
	fmt.Println("删除后的整数切片:", sliceV3)

	sliceV4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("原始整数切片:", sliceV4, "容量:", cap(sliceV4))
	for i := 0; i < 8; i++ {
		sliceV4 = deleteAtIndexV4(sliceV4, 0)
		fmt.Println("删除后的整数切片:", sliceV4, "容量:", cap(sliceV4))
	}
}

// 删除切片中指定索引的元素
func deleteAtIndexV1(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s // 索引超出范围，返回原切片
	}
	return append(s[:index], s[index+1:]...)
}

// 删除切片中指定索引的元素
func deleteAtIndexV2(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s // 索引超出范围，返回原切片
	}
	copy(s[index:], s[index+1:])
	return s[:len(s)-1]
}

func deleteAtIndexV3[T any](s []T, index int) []T {
	if index < 0 || index >= len(s) {
		return s // 索引超出范围，返回原切片
	}
	copy(s[index:], s[index+1:])
	return s[:len(s)-1]
}

const shrinkThreshold = 0.25

// 删除切片中指定索引的元素并可能缩容
func deleteAtIndexV4[T any](s []T, index int) []T {
	if index < 0 || index >= len(s) {
		return s // 索引超出范围，返回原切片
	}
	copy(s[index:], s[index+1:])
	s = s[:len(s)-1]

	// 缩容机制
	if float64(len(s))/float64(cap(s)) < shrinkThreshold {
		newSlice := make([]T, len(s), len(s)*2)
		copy(newSlice, s)
		return newSlice
	}

	return s
}
