// In this problem we have implemented stack using linked list which is used by producer to  produce the data to bufer and   consumed by consumer using non-concurrent way.
//In stack we usually add new node at tail and remove it from head
package main

//Here we have defined all the packages we have used in our code
//fmt function we have used to print the data on console, Package strconv implements conversions to and from string representations of basic data types,time we have used to use functions which calculate time of execution of whole program,

import (
  "fmt"
	"strconv"
	"time"
  
)
//Here we have defined the structure of our node.Index is the index has values of the producer,producer has id for producer number,next and prev are pointers of node which play a key role in making linked list.


type Node struct {
  index           int   
	producer        string
  next,prev *Node
  
}
//This struct type stores the head and tail pointer of linked list
type List1 struct {
    name       string
    head, tail *Node
}

func main() {
	
 var n,size int

  fmt.Println("Enter number of producer and consumer")
  fmt.Scan(&n)
  fmt.Println("you have entered", n," as num of producer and consumer")
  fmt.Println("Enter the buffer size")
  fmt.Scan(&size)
  fmt.Println("you have entered", size,"buffer size")

  //creating buffer
  p :=  createBuffer("buffer")
  //j is the variable to keep track of size of buffer
  j:=0
  // here start we capture time stamp of start of producer consumer problem
  start := time.Now()
  
	for i := 1; i < n; i++ {
    p.producer(size,i,&j)
		time.Sleep(1 * time.Second)
    p.consumer(size,i,&j)
		time.Sleep(1 * time.Second)
	}

	elapsed := time.Since(start)
  fmt.Println(elapsed)
}


func (p *List1) producer(size int,id int,j *int) {
	
	for i := 0; i < 10; i++ {
    //defining producer name
    producer_name := "producer"
		producer_name += strconv.Itoa(id)
    //if the buffer size is full producer should wait in infinite loop for any producer to consume and make some size free.
		
    for *j>=size {
      fmt.Println("*************buffer full*********")
    }
    
    
    //call push function
    p.push(i, producer_name)
    //increment the buffer filled size
     *j=*j+1;
		fmt.Println("This is number", i, "--producer:", id)
    
		time.Sleep(1 * time.Second)

		
	}
  

}
func (p *List1) consumer(size int,id int,j *int) {
	for i := 0; i < 10; i++ {
        //if there is nothing to consume for consumer to consume we should wait in infinite loop
		for *j <= 0{
       fmt.Println("*************nothing to consume*********")
    }
    
		fmt.Println("Consumer--", id, " --buy the number--", p.tail.producer,"consuming number",p.tail.index)
    
    //pop the value
    p.pop()

    // decrement the pointer when you are done
    *j=*j-1
    time.Sleep(1 * time.Second)
  }
  
}

func (p *List1) push(index int,producer_name string) {
//Make simple null node
s := &Node{
		index:  index,
    producer        :producer_name,
		
	}

//if p.head is null it means we are adding first node
  if p.head == nil {
    p.head=s
    s.next=s
	}else {
    //if tree is not null then we put s node to next pointer of last node
		currentNode := p.tail
		currentNode.next = s
		s.prev = p.tail
    
	}
   //now we increment the tail to new added node
	p.tail = s
  
	
}

//pop function
func (p *List1) pop() {
//if linked list head is nil that means that we dont have anything to delete
//check p.tail is nil or not if nil we dont have any stack
if(p.tail!=nil){
  //see if there is only one node in stack
  if(p.tail==p.head){
   p.head=p.head.prev
  }else {
    //if not then delink the second last node
   p.tail.prev.next = p.head.prev
   p.tail=p.tail.prev
  }
}
  
}
func createBuffer(name string) *List1 {
	return &List1{
		name: name,
	}
}