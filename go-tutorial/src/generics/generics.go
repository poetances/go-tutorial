package generics

import (
	"context"
	"fmt"
)

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

// 3.泛型接口
type Comparable[T any] interface {
	Compare(other T) int
}

type Container[T any] interface {
	Add(item T)
	Remove(item T) bool
	Contains(item T) bool
}

type PairInterface[K, V any] interface {
	Key() K
	Value() V
	SetKey(key K)
	SetValue(value V)
}

// 4.类型约束
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
type MyInt int     // 底层类型是 int
type YourInt = int // 底层类型是 int

// 5.多参数类型约束
func Swap[T1, T2 any](a T1, b T2) (T2, T1) {
	return b, a
}

type Pair[K, V any] struct {
	key   K
	value V
}

func (p Pair[K, V]) Key() K {
	return p.key
}
func (p Pair[K, V]) Value() V {
	return p.value
}

type AdderFunc[T any] func(a, b T) T

func AddTest[T int](a AdderFunc[T]) T {
	return a(1, 2)
}

//////////////////////////////////////////////////////////////////
// Handler实现了一个基于Go泛型的异步队列处理器，它是一个通用的生产者-消费者模式实现，可以处理任意类型的数据项。
/////////////////////////////////////////////////////////////////
type Handle[T any] func(item T)
type Handler[T any] struct {
	cancel context.CancelFunc
	itemChan   chan T
}

func NewHandler[T any]() *Handler[T] {
	return &Handler[T]{}
}
func (h *Handler[T]) Start(size int, handle Handle[T]) {
	h.Stop()
	ctx, cancel := context.WithCancel(context.Background())
	h.itemChan = make(chan T, size)
	h.cancel = cancel
	go func() {
		for {
			select {
			case item := <-h.itemChan:
				handle(item)
			case <-ctx.Done():
				return
			}
		}
	}()
}
func (h *Handler[T]) Stop() {
	if h.cancel != nil {
		h.cancel()
		h.cancel = nil
		h.itemChan = nil
	}
}
func (h *Handler[T]) Push(item T) bool{
	if h.itemChan == nil {
		return false
	}
	select {
	case h.itemChan <- item:
		return true
	default:
		return false
	}
}