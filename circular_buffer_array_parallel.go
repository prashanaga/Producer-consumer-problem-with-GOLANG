package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var front = -1
var rear = -1
var queue [5]int
var limit = 5
var m int
var mutex = &sync.Mutex{}
var wg = &sync.WaitGroup{}

//function to insert elements in the empty queue

func insert_empty(limit int,i int) {

	front = 0
	rear = 0
	m = rand.Intn(100)
	queue[rear] = m
	fmt.Println("---Producer---",i,"---Element produced-- : ", m)
	fmt.Println("Queue is :", queue)

}

//function to insert elements

func insert(limit int,i int) {

	m = rand.Intn(100)
	rear = rear + 1
	queue[rear] = m
	fmt.Println("---Producer---",i,"---Element produced-- : ", m)
	fmt.Println("Queue is :", queue)

}

//function to insert elements in the queue when the rear is at max limit and front not equal to 0

func insert_special(limit int,i int) {

	if rear == limit-1 {

		rear = 0
		m = rand.Intn(100)
		queue[rear] = m

	} else {

		m = rand.Intn(100)
		for i := rear + 1; i <= front-1; i++ {

			queue[i] = m
			rear = i

		}

	}
	fmt.Println("---Producer---",i,"---Element produced-- : ", m)
	fmt.Println("Queue is :", queue)

	//time.Sleep(2 * time.Millisecond)
}

// function to check the queue is FULL,EMPTY  and insertion of elements

func enqueue(limit int,i int) {
	mutex.Lock()
	defer mutex.Unlock()
	defer wg.Done()
	if (front == rear+1) || (front == 0 && rear == limit-1) {
		fmt.Println("Can not produce..Queue is FULL !!")

	} else if rear == limit-1 && front != 0 {
		insert_special(limit,i)

	} else if rear == -1 && front == -1 {
		insert_empty(limit,i)

	} else {
		insert(limit,i)
	}

}

// function to check the queue is empty or else delete elements

func dequeue(limit int,i int) {

	var num_deleted int
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()
	if front == -1 {
		fmt.Println("Can not consume..Queue is Empty !!")

	} else if front == rear {

		num_deleted = queue[front]
		fmt.Println("---consumer---",i,"---consumed element is--- :", num_deleted)
		queue[front] = 0
		front = -1
		rear = -1
		fmt.Println("Queue is :", queue)

	} else {

		num_deleted = queue[front]
		fmt.Println("---consumer---",i,"---consumed element is--- :", queue[front])
		queue[front] = 0
		if front == limit-1 {
			front = 0
		} else {
			front = front + 1

		}

		fmt.Println("Queue is :", queue)

	}

}

// function to display elements

func display() {

	fmt.Println("queue is :", queue)

}

func main() {

	var prd_num, cons_num int
	start := time.Now()
	fmt.Println("Enter the No.of Producers :")
	fmt.Scan(&prd_num)
	fmt.Println("Enter the No.of Consumers :")
	fmt.Scan(&cons_num)

	for i := 0; i < prd_num; i++ {
		wg.Add(1)
		go enqueue(limit,i)
	}
	for i := 0; i < cons_num; i++ {
		wg.Add(1)
		go dequeue(limit,i)
	}

	wg.Wait()
	fmt.Println("Program Done")
	elapsed := time.Since(start)
	fmt.Println("Time execution", elapsed)
}
