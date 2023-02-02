package utils

import (
	. "fmt"
)

func init() {}

type numeric interface {
	~int | byte | ~uint32 | float32 | ~float64
}

type List[T numeric] struct {
	arr []T
	len int
	cap int
}

func (a *List[T]) Append(element T) {
	if a.cap == a.len {
		newArr := a.Extension()
		//	   将旧值移动到新的数组中
		for i := 0; i < a.len; i++ {
			newArr.arr[i] = a.arr[i]
		}
		//       将新数组赋值给原数组
		a.arr = newArr.arr
		//       改变数组的map
		a.cap = newArr.cap
	}
	a.arr[a.len] = element
	//     原数组长度+1
	a.len = a.len + 1
}

func (a *List[T]) Extension() *List[T] {
	var newCap int
	if a.cap <= 10 {
		newCap = 3 * a.len
	} else {
		newCap = 2 * a.len
	}
	// 新数组的长度为newCap
	newArr := make([]T, newCap, newCap)
	// arr的长度应该也是newCap
	a.arr = newArr
	a.cap = newCap
	return a
}

func MakeList[T numeric](len, cap int) *List[T] {
	list := new(List[T])
	if len > cap {
		Println("invalid，len large than cap")
	}
	arr := make([]T, len, cap)
	list.arr = arr
	list.len = 0
	list.cap = cap
	return list
}
