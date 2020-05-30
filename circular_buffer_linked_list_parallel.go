// In this problem we have implemented circular buffer using linked list which is used by producer to continuously produce the data to bufer and continuously it is consumed by consumer.
//In circular buffer we usually have queue implementation wherein we add new node at tail and remove it from head

package main
//Here we have defined all the packages we have used in our code
//fmt function we have used to print the data on console, Package strconv implements conversions to and from string representations of basic data types,time we have used to use functions which calculate time of execution of whole program,Package sync provides basic synchronization primitives such as mutual exclusion locks and waitgroup functionality

import (
  "fmt"
	"strconv"
	"time"
	"sync"
  
)
//Here we have defined the structure of our node.Index is the index has values of the producer,producer has id for producer number,next and prev are pointers of node which play a key role in making linked list.

type Node struct {
  index           int    //production id index
	 //name of production
	producer        string
  next,prev *Node
  
}

//This struct type stores the head and tail pointer of linked list
type List1 struct {
    name       string
    head, tail *Node
}
//here is the start of main function,when program starts the compiler first strt implementing main function
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

  //A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished like here we used  wg.Add(1) and wg.wait() to let all go routines finish

  var wg sync.WaitGroup
  //A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex. In this program we used mutex to lock shared variable.
  var m sync.Mutex
	for i := 1; i < n; i++ {
    wg.Add(1)
    //producing go routines for producer using call by reference
		go p.producer(size,i,&m,&wg,&j)

		time.Sleep(1 * time.Second)
    wg.Add(1)
    //producing go routines for consumer using call by reference
		go p.consumer(size,i,&m,&wg,&j)
		time.Sleep(1 * time.Second)
	}

	// Just delay the end of main
	wg.Wait()
  elapsed := time.Since(start)
  fmt.Println(elapsed)
	

 	  
}


func (p *List1) producer(size int,id int,mu *sync.Mutex ,wg *sync.WaitGroup,j *int) {
	
  for i := 0; i < 10; i++ {
    
    //defining producer name
    producer_name := "producer"
		producer_name += strconv.Itoa(id)
		//if the buffer size is full producer should wait in infinite loop for any producer to consume and make some size free.

    for *j>=size {
      fmt.Println("*************buffer full*********")
    }


//Now since we are going to add values to buffer we take lock

    mu.Lock()
    
    // call add function
    p.add(i, producer_name)
    //increment the buffer size
    *j=*j+1;
    //
		fmt.Println("This is number", i, "--producer:", id)

    //remove the lock
    mu.Unlock()
		
		time.Sleep(1 * time.Second)

		
	}
  //after processing of this thread we are telling wait group this this g routine is done you canmark as done
  wg.Done()

}

//function for consumer 
func (p *List1) consumer(size int,id int,mu *sync.Mutex ,wg *sync.WaitGroup,j *int) {
	for i := 0; i < 10; i++ {
        
		//if there is nothing to consume for consumer to consume we should wait in infinite loop
    for *j <= 0{
       fmt.Println("*************nothing to consume*********")
    }
    //now take lock because we are going to delete the node and we don't want any other thread use this share resource

    mu.Lock()

    fmt.Println("Consumer--", id, " --buy the number--", p.head.producer,"consuming number",p.head.index)


   
    p.delete()
    // decrement the pointer when you are done
    *j=*j-1;
    mu.Unlock()
		time.Sleep(1 * time.Second)

	}
  //after processing of this thread we are telling wait group this this g routine is done you canmark as done
  wg.Done()
}




func (p *List1) add(index int,producer_name string) {
//Make simple null node
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
        fmt.Println ("Linked list is empty"); 
        
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