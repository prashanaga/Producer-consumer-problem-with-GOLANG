package main

// In this problem we have implemented circular buffer using linked list which is used by producer to  produce the data to bufer and it is consumed by consumer not in parallel way.
//In circular buffer we usually have queue implementation wherein we add new node at tail and remove it from head
import (
	"fmt"
	"strconv"
	"time"
)

type Node struct {
  index           int    //production id index
  producer        string
  next,prev *Node
}

type List1 struct {
    name       string
    head, tail *Node
}

func main() {
	//n is number of producer and consumer defined
  //size is the size of buffer
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
    
    producer_name := "producer"
		producer_name += strconv.Itoa(id)
		
    //if the buffer size is full producer should wait in infinite loop for any producer to consume and make some size free.
    for *j>=size {
      fmt.Println("*************buffer full*********")
    }
   // call add function
    p.add(i, producer_name)
    //increment the buffer size
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
   
    
		fmt.Println("Consumer--", id, " --buy the number--", p.head.producer,"consuming number",p.head.index)
    // decrement the pointer when you are done
    *j=*j-1;
    p.delete()
    
		time.Sleep(1 * time.Second)

	}
 
}


func (p *List1) add(index int,producer_name string) {

s := &Node{
		index:  index,
    producer        :producer_name,
		
	}
//if p.head is null it means we are adding first node
    if (p.head == nil) {
        p.head = s; 
    }else{
      //if tree is not null then we put s node to next pointer of last node
        p.tail.next = s; 
    }
   //now we increment the tail to new added node
    p.tail = s; 
     //we add next pointer of last node to head of linked list
    p.tail.next =p.head;
  
	
}

//delete function
func (p *List1) delete() {
//if linked list head is nil that means that we dont have anything to delete

if (p.head == nil){ 
        fmt.Println ("Queue is empty"); 
        
    } 
  
    // If this is the last node to be deleted 
     // Value to be deleted  
    if (p.head == p.tail) { 
        p.head = nil; 
        p.tail = nil; 
    } else{ 
       //else make head next node of linked list 
        p.head = p.head.next; 
        // change the tail pointer of linked list to new head
        p.tail.next= p.head; 
         
    } 


  

}

func createBuffer(name string) *List1 {
	return &List1{
		name: name,
	}
}