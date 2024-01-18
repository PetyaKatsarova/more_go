package main

import "fmt"

func main() {
	finished := make(chan bool)
	names := []string{"Pip"}

	go func() {
		names = append(names, "Electric")
		names = append(names, "Tinkie Winkie")
		finished <- true
	}()

	for _, name := range names {
		fmt.Println(name)
	}
	<- finished 
}