package generics

import "fmt"

/*
一、什么是泛型？
泛型是一种编程特性，允许我们在编写代码时使用类型参数，而不是具体类型。这样可以让函数和数据结构适用于多种类型，提高代码的复用性。
*/

// 1.泛型基本语法
func PrintSlice[T any](s []T) {
	for i, v := range s {
		fmt.Printf("索引 %d: 值 %v (类型 %T)\n", i, v, v)
	}
}

// 泛型约束
// 类型参数约束语法详解
// T          - 类型参数名称
// int | float64 - 类型约束，可以是 int 或 float64
// any        - 预定义的约束，表示任何类型（等同于 interface{}）
func Add[T int | float64](a, b T) T {
	return a + b
}

// 2.泛型结构体
type Stack[T any] struct {
	elements []T
}
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index]
	return element, true
}
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// 3.泛型结构
type Comparable[T any] interface {
	Compare(other T) int
}
// 类型约束
type Number interface {
	int | int32 | int64 | float32 | float64
}
func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}
type Addable interface {
    ~int | ~float32 | ~float64 // ~ 表示底层类型, 包括自定义类型
}
// 底层类型相同：两个类型即使名称不同，只要底层类型一致即视为相同。比如：
type MyInt int    // 底层类型是 int
type YourInt = int // 底层类型是 int
