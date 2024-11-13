package main

import (
	"fmt"
)

type MinStack struct {
	Stack []int
	Min []int
	Tops int
}

func (this *MinStack) Push(val int) {
	if this.Tops == -1 {
		this.Min = append(this.Min, val)
	} else {
		this.Min = append(this.Min, Min(val, this.Min[this.Tops]))
	}
	this.Stack = append(this.Stack, val)
	this.Tops++
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}, -1}
}

func (this *MinStack) Pop()  {
	this.Stack = this.Stack[:this.Tops]
	this.Min = this.Min[:this.Tops]
	this.Tops--   
}

func (this *MinStack) Top() int {
	return this.Stack[this.Tops]
}

func (this *MinStack) GetMin() int {
	return this.Min[this.Tops]
}


func Min(a,b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	stack := Constructor()
	stack.Push(100)
	stack.Push(50)
	stack.Push(300)
	stack.Pop()
	stack.Push(200)
	fmt.Println(stack)
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
}
