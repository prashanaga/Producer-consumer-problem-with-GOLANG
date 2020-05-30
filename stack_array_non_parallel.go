package main

import (
	"fmt"
	"math/rand"
	"time"
)

var inpt string

var top = -1
var le = -1
var n int
var stack []int
var limit int

// Code to push element to stack

func push(limit int) {

	var m int

	if top >= limit-1 {

		fmt.Println("stack is full")

	} else {

		m = rand.Intn(100)
		top += 1
		fmt.Println("pushed element :", m)
		stack = append(stack, m)

	}
	display()

}

// Code to pop element from stack

func pop() {

	if top <= -1 {

		fmt.Println("Can not POP; Stack is EMPTY")
	} else {

		fmt.Println("Poped element :", stack[top])

		stack = stack[:top]
		top -= 1

	}

	display()
}

// Code to display elements

func display() {

	fmt.Println(stack)

}

func main() {

	var prod_cons_num int
	start := time.Now()

	fmt.Println("Enter limit of stack :")
	fmt.Scan(&limit)
	fmt.Println("Enter the no of producers and consumers :")
	fmt.Scan(&prod_cons_num)

	for i := 0; i < prod_cons_num; i++ {
		push(limit)

	}
	for i := 0; i < prod_cons_num; i++ {
		pop()

	}

	elapsed := time.Since(start)
	fmt.Println("Time execution", elapsed)

}
