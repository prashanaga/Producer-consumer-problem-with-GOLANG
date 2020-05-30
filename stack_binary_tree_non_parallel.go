package main

import (
	"fmt"
	"math/rand"
	"time"
)

var stack []int

type node struct {
	val   int
	left  *node
	right *node
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type rootnode struct {
	root *node
}

var btree = &rootnode{}

// function to insert elements
func (t *rootnode) insert(data int) *rootnode {

	if t.root == nil {

		t.root = &node{val: data, left: nil, right: nil}
		stack = append(stack, data)
	

	} else {
		t.root.insert(data)
	}
	fmt.Println("Inorder traversal is :")
	    inorder_traverse(btree.root)
	    fmt.Println(" ")
	    fmt.Println("stack is", stack)
	return t
}

func (d *node) insert(data int) {

	if d == nil {
		return

	} else if data == d.val {

		fmt.Printf(" Sorry, element already existing in the tree %v \n", data)
	} else if data <= d.val {

		if d.left == nil {
			d.left = &node{val: data, left: nil, right: nil}
			stack = append(stack, data)

		} else {

			d.left.insert(data)
		}

	} else {

		if d.right == nil {
			d.right = &node{val: data, left: nil, right: nil}
			stack = append(stack, data)
		} else {

			d.right.insert(data)
		}

	}
	

}
// function to display elements
func inorder_traverse(d *node) {
	if d == nil {
		return

	} else {
		inorder_traverse(d.left)
		fmt.Printf("%v ", d.val)
		inorder_traverse(d.right)
	}

}
// function to pop elements
func delete_element(d *node, prevnode *node, poped_element int) {
	temproot := d

	if temproot.val > poped_element { // i used swap technique
		prevnode = temproot
		temproot = prevnode.left
		delete_element(temproot, prevnode, poped_element)

	} else if temproot.val < poped_element {
		prevnode = temproot
		temproot = prevnode.right
		delete_element(temproot, prevnode, poped_element)

	} else {
		if temproot.val < prevnode.val {
			fmt.Printf("consumed element is %v \n", temproot.val)
			prevnode.left = nil
		} else if temproot.val > prevnode.val {
			fmt.Printf("consumed element is %v \n", temproot.val)
			prevnode.right = nil
		} else { //only root node is remaining
			fmt.Printf("consumed element is %v \n", temproot.val)
			prevnode = nil
		}
	}
}

func main() {

	var limit int
	start := time.Now()
	fmt.Println("Enter no of producers  :")
	fmt.Scan(&limit)

	for i := 0; i < limit; i++ {
		producedelement := rand.Intn(100)
		fmt.Println("produced element", producedelement)
		btree.insert(producedelement)

	}

	for i := 0; i < 10; i++ {
		if len(stack) > 0 {
			element := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			delete_element(btree.root, btree.root, element)
		}

	}
	elapsed := time.Since(start)
	fmt.Println("Time execution", elapsed)
}
