// In this problem we have implemented stack using avl which is used by producer to  produce the data to bufer and   consumed by consumer using non-concurrent way.
//In stack we usually add new node at tail and remove it from head
package main

import (
	
	"fmt"
  "strconv"
	"time"
  "math/rand"
	
)

type node struct {
  Key            int
	Height         int
  lchild,rchild *node
  
}

func main() {
	var n,size int

  fmt.Println("Enter number of producer and consumer")
  fmt.Scan(&n)
  fmt.Println("you have entered", n," as num of producer and consumer")
  fmt.Println("Enter the buffer size")
  fmt.Scan(&size)
  fmt.Println("you have entered", size,"buffer size")
//making root node
  var root *node
  //j is the variable to keep track of size of buffer
  j:=0
  start := time.Now()
 
	for i := 1; i < n; i++ {
    

		producer(size,i,&j,root)
		
		time.Sleep(1 * time.Second)
	}

	
  elapsed := time.Since(start)
  fmt.Println(elapsed)
	
	
 	  
}


func  producer(size int,id int,j *int,root *node)*node {
	
  
	for i := 0; i < 10; i++ {
    
    
		name_product := rand.Intn(100)
		producer_name := "producer"
		producer_name += strconv.Itoa(id)
		
    for *j>=size {
      fmt.Println("*************buffer full*********")
    }
   
    
    *j=*j+1;
    root=insert(root, name_product)
		fmt.Println("This is number", i, "--producer:", id,"product name",name_product)
   
    time.Sleep(1 * time.Second)

		
	}
   
  
  return root

}
func  consumer(size int,id int,j *int,root *node)*node {
	for i := 0; i < 10; i++ {
        
		//buy := <-c
    for *j <= 0{
       fmt.Println("*************nothing to consume*********")
    }
   
		//fmt.Println("Consumer--", id, " --buy the number--", //p.head.producer,"consuming number",p.head.index)
    *j=*j-1;
    ///delete()
   
		time.Sleep(1 * time.Second)

	}
 
  return root
}

//If a tree becomes unbalanced, when a node is inserted into the right subtree of the right subtree, then we perform a single left rotation
func leftRotate(root *node) *node {
	node := root.rchild
	// fmt.Println(node.Key)
	root.rchild = node.lchild
	node.lchild = root

	root.Height = max(height(root.lchild), height(root.rchild)) + 1
	node.Height = max(height(node.rchild), height(node.lchild)) + 1
	return node
}
//its a type of double rotation,its a combination of left rotation followed byright rotation
func leftRigthRotate(root *node) *node {
	root.lchild = leftRotate(root.lchild)
	root = rightRotate(root)
	return  root
}
//if a node inserted in left subtree  and it becomes unbalanced then we do right rotation,its a type of single rotation
func rightRotate(root *node) *node {
	node := root.lchild
	root.lchild = node.rchild
	node.rchild = root
	root.Height = max(height(root.lchild), height(root.rchild)) + 1
	node.Height = max(height(node.lchild), height(node.rchild)) + 1
	return node
}
// A left-right rotation is a combination of left rotation followed by right rotation.
func rightLeftRotate(root *node) *node {
	root.rchild = rightRotate(root.rchild)
	root = leftRotate(root)
	return  root
}
//function to find the height of tree
func height(root *node) int {
	if root != nil {
		return root.Height
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(root *node, key int) *node {
  //if root is null make your root
	if root == nil {
		root = &node{key, 0, nil, nil}
		root.Height = max(height(root.lchild), height(root.rchild)) + 1
		return root
	} 
// check the value to insert if value to be insertes < root key insert it to left child
	if key < root.Key {
		root.lchild = insert(root.lchild, key)
    //this piece of code is for balancing the tree
		if height(root.lchild)-height(root.rchild) >= 2 {
			if key < root.lchild.Key {
				root = rightRotate(root) 
			} else {
				root = leftRigthRotate(root)
			}
		}
	} 
//if key is more than root.key we try to insert on right side
	if key > root.Key {
		root.rchild = insert(root.rchild, key)
    //this piece of code is for checing whether tree is balanced or not
		if height(root.rchild)-height(root.lchild) >= 2 {
			if key > root.rchild.Key {

				root = leftRotate(root) 
			} else {
				root = rightLeftRotate(root) 
			}
		}
	}

	root.Height = max(height(root.lchild), height(root.rchild)) + 1
	return root
}



type action func(node *node)

func inOrder(root *node, action action) {
	if root == nil {
		return
	}

	inOrder(root.lchild, action)
	action(root)
	inOrder(root.rchild, action)
}