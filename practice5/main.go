package main

import (
	"fmt"
)

func main() {
	doSomething()

	sum := addValues(5, 8)

	fmt.Println(sum)

	multisum, multicount := addAllValues(4, 3, 5)
	fmt.Println(multisum)
	fmt.Println(multicount)

	//custom types/structs allow for methods

	poodle := Dog{"poodle", 15, "bark"}

	poodle.Speak()
	poodle.SpeakThrice()
	poodle.Speak()

}

func doSomething() {
	fmt.Println("doing something")
}

// both parameters will beint, because of the order
func addValues(value1, value2 int) int {
	return value1 + value2
}

// pass multiple vals of same type
func addAllValues(values ...int) (int, int) { //two return types mean we return two values
	total := 0
	for _, val := range values {
		total += val
	}
	return total, len(values)
}

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// this first set of parenthesis is a receiver, of type Dog (our custom type)
// this seems to be how we incorporate static methods into Go
//for these functions we have to identify what it does with a comment, kinda like an annotation

//the scope for these methods is dynamic, so as we update sound, Speak() will change what the output is

// Speak is how the dog speaks
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d Dog) SpeakThrice() {
	//this is a copy of the dog object, so it runs three times every time this is called
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
