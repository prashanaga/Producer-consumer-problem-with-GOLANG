package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var inpt string

var top = -1
var le = -1
var n int
var stack []int
var limit int
var mutex = &sync.Mutex{}
var wg = &sync.WaitGroup{}

// Code to push element to stack

func push(limit int,i int) {

	defer wg.Done()
	var m int

	if top >= limit-1 {

		fmt.Println("stack is full")

	} else {
		mutex.Lock()
		m = rand.Intn(100)
		top += 1
		fmt.Println("----Producer---",i,"--pushed element--:", m)
		stack = append(stack, m)
		display()
		mutex.Unlock()

	}
	
	time.Sleep(2 * time.Millisecond)

}

// Code to pop element from stack

func pop(i int) {
	
	defer wg.Done()
	if top == -1 {
		fmt.Println("Can not Consume; Stack is EMPTY")
	} else {
	    mutex.Lock()
		fmt.Println("-----consumer----",i,"---Poped element----:", stack[top])

		stack = stack[:top]
		top -= 1
		display()
        mutex.Unlock()
		
	}
	
	time.Sleep(2 * time.Millisecond)
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
		wg.Add(1)
		go push(limit,i)
		time.Sleep(1 * time.Millisecond)
		wg.Add(1)
		go pop(i)

	}

	wg.Wait()
	fmt.Println("Program Done")
	elapsed := time.Since(start)
	fmt.Println("Time execution", elapsed)

}
