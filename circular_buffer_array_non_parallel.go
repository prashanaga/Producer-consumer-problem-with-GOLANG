package main

import (
	"fmt"
	"math/rand"
	"time"
)

var front = -1
var rear = -1
var queue [5]int
var limit = 5
var m int

//function to insert elements in the empty queue

func insert_empty(limit int) {
	front = 0
	rear = 0
	m = rand.Intn(100)
	queue[rear] = m
	display()
}

//function to insert elements

func insert(limit int) {

	m = rand.Intn(100)
	rear = rear + 1
	queue[rear] = m

	display()
}

//function to insert elements in the queue when the rear is at max limit and front not equal to 0

func insert_special(limit int) {
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

	display()
}

// function to check the queue is FULL,EMPTY  and insertion of elements

func enqueue(limit int) {

	if (front == rear+1) || (front == 0 && rear == limit-1) {

		fmt.Println("Queue is FULL !!")
	} else if rear == limit-1 && front != 0 {

		insert_special(limit)

	} else if rear == -1 && front == -1 {

		insert_empty(limit)

	} else {
		insert(limit)
	}

}

// function to check the queue is empty or else delete elements

func dequeue(limit int) {
	if front == -1 {

		fmt.Println("Queue is Empty !!")
	} else {
		delete(limit)

	}

}

// function to delete elements from queue

func delete(limit int) {

	var num_deleted int

	num_deleted = queue[front]

	if front == rear {

		fmt.Println("Deleted element is :", num_deleted)
		queue[front] = 0
		front = -1
		rear = -1
		fmt.Println("Queue is :", queue)

	} else {

		fmt.Println("Deleted element is :", queue[front])
		queue[front] = 0
		if front == limit-1 {
			front = 0
		} else {
			front = front + 1

		}

		display()

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
		enqueue(limit)
	}

	for i := 0; i < cons_num; i++ {
		dequeue(limit)
	}

	//	dequeue(limit)
	//	enqueue(limit)
	
	elapsed := time.Since(start)
	fmt.Println("Time execution", elapsed)
}
