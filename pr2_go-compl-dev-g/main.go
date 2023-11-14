package main

import "fmt"

// func main() {
// 	// cards := newDeck()
// 	// cards.saveToFile("my_cards")
// 	//cards.print()
// 	// print(newDeck())
// 	// fmt.Println(newDEckFromFile("my_cards"))

// 	// hand, remainingGards := deal(cards, 5)
// 	// fmt.Println("------------------------hand cards----------------------------------- ")
// 	// hand.print()
// 	// fmt.Println("--------------------- remainining cards -------------------------")
// 	// remainingGards.print()

// 	// ------------ type conversion ----------------
//     // []byte("Hi there!") where []byte is the type we want and 'hi there' is the value we have: like type casting
// 	greeting := "hi there"
// 	fmt.Println([]byte(greeting))

// 	// fmt.Println(cards.toString())
// 	// cards := newDeck()
// 	// cards.shuffle()
// 	// cards.print()

// 	p := Person {
// 		fname: "Pip",
// 		lname: "K",
// 		contact: ContactInfo{
// 			address: "uweg",
// 			code: "1231tk",
// 		},
// 	}

// 	// p.fname = "jimmy"
// 	pointerToP := &p
// 	pointerToP.updateFname("jimmy") // updateFname() takes a pointer to *Person!
// 	p.updateFname("lala") // shortcut because in updateFname() it takes a pointer to person
// 	fmt.Printf("%+v\n", p)
// 	p.updateFname("petka")
// 	fmt.Printf("%+v\n", p)
// }

// func (p *Person) updateFname(fname string) {
// 	// p.fname = fname
// 	(*p).fname = fname // *p we want to manipulate the val the pointer is referencing to
// }

// type ContactInfo struct {
// 	address string
// 	code string
// }

// type Person struct {
// 	fname string
// 	lname string
// 	contact ContactInfo
// }


// func main() {
// 	mysl := []string{"hi", "schatje"}
// 	updateSl(mysl)
// 	fmt.Println(mysl)
// }

// func updateSl(s []string) {
// 	s[0] = "bye"
// }

// ------------ MAPS --------------------------------

// func main() {
// 	colors := map[string] string { // key is string, val is string
// 		 "red": "#ff0000",
// 		 "green": "#4bf745",
// 		"white": "#ffffff",
// 	}
// 	printMap(colors)
// 	// fmt.Println(colors)
// 	// delete(colors, "red")
// 	// fmt.Println(colors)
// }

// func printMap(c map[string] string) {
// 	for color, hex := range c {
// 		fmt.Println("hex code for", color, "is", hex)
// 	}
// }

// -------------------- 	INTERFACES ---------------------------------------

func main() {

}

type bot interface {
	getGreeting() string
}