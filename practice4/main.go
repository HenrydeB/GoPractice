package main

import (
	"fmt"
)

func main() {

	/* theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "less than zero"
	} else if theAnswer == 0 {
		result = "the answer is zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)

	if x := -42; x < 0 { //you can declare variables in the if statement
		result := "less than zero"
	} else {
		result := "greater than or equal to zero"
	}

	fmt.Println(result) */

	// SWITCH statements
	/* rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)

	var result string
	switch dow {
	case 1:
		result = "its sunday"
	case 2:
		result = "its monday"
	default:
		result = "it's some other day"
	}

	//note that we don't need a break statement

	//if we add "fallthrough, it'll move to the next statement"
	fmt.Println(result) */

	//FOR loops

	colors := []string{"red", "green", "blue"}

	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	//this can be written more concicely with 'range'

	for i := range colors {
		//this sets the index to each position in the loop

		fmt.Println(colors[i])
	}

	//we can also do a for each, using two values: index, and current value
	// if we don't need the index, we can use _

	for _, color := range colors {
		fmt.Println(color)
	}

	//next is essentially a while loop
	value := 1
	for value < 10 {
		fmt.Println(value)
		value++
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)

		if sum > 200 {
			goto theEnd //jumps to the goto location
		}
	}

theEnd:
	fmt.Println("End of Program")
}
