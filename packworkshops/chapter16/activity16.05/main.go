package main

import (
	"fmt"
	"log"
)

func greet(ch chan string) {
	msg := <-ch //receives a message from the channel ch and stores it in the variable msg. This is a blocking operation, 
	//meaning if there's no data available in the channel, it will wait until there is.
	ch <- fmt.Sprintf("Thanks for %s", msg) // send msg to the ch
	ch <- "Hello Pip"
}

func main() {
	ch := make(chan string)

	go greet(ch) // if you run greet(ch) without the go keyword, it will not work as expected, and it will block the execution
	// of the main function until greet finishes its execution. it is explained at the bottom

	ch <- "Hello John"

	log.Println(<-ch)
	log.Println(<-ch)
}

/*
When you call greet(ch) without the go keyword, it runs greet as a regular function call, not as a separate goroutine. 
This means that the greet function will execute sequentially within the main function, and any operations inside greet that block
 (such as channel reads or writes) will block the entire program's execution.
Blocking Behavior: Inside the greet function, you have channel operations like <-ch and ch <-. These operations will block
 until there is a sender or receiver on the other side of the channel. If there is no other goroutine reading from or writing
  to the channel, the program will hang, and you won't see further progress.
*/