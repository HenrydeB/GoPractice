package main

import "fmt"

const num int = 54

func main() {
	//points to a memory address that has nothing set there
	/* anInt := 32
	var p = &anInt
	fmt.Println(*p) */

	//Ordered values in array

	/* var colors [3]string
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"

	fmt.Println(colors)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers) */

	//slices (resizable, sorted easily)
	// initialized as an array without a predetermined size

	/* var things []string
	things = append(things, "stuff")
	fmt.Println(things)

	stuff := make([]int, 5, 5)

	stuff[0] = 3
	stuff[1] = 2
	stuff[2] = 1
	stuff[3] = 5
	stuff[4] = 4

	fmt.Println(stuff)
	stuff = append(stuff, 6)
	stuff = append(stuff, 7)
	fmt.Println(stuff)
	sort.Ints(stuff)
	fmt.Println(stuff) */

	//store unordered values in maps

	/* 	states := make(map[string]string)
	   	fmt.Println(states)
	   	states["WA"] = "Washington"
	   	states["IL"] = "Illinois"
	   	states["CA"] = "California"

	   	fmt.Println(states)

	   	california := states["CA"]
	   	fmt.Println(california)

	   	delete(states, "OR")
	   	fmt.Println(states)
	   	states["NY"] = "New York"
	   	fmt.Println(states)

	   	for k, v := range states {
	   		fmt.Printf("%v, %v\n", k, v) //doing this as is will print out the keys in any order, not specific
	   	}

	   	keys := make([]string, len(states))

	   	i := 0
	   	for k := range states {
	   		keys[i] = k
	   		i++
	   	}

	   	fmt.Println(keys)
	   	sort.Strings(keys)
	   	fmt.Println(keys)

	   	for j := range keys {
	   		fmt.Println(states[keys[j]])
	   	} */

	//STRUCTS
	//no inheritence, Structs are custom types
	poodle := Dog{"poodle", 50}
	fmt.Println(poodle)

	fmt.Printf("%+v\n", poodle) //allows us to prints the fields and their values
	//can also do this with
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

}

// upper case means it is exported
type Dog struct {
	Breed  string
	Weight int
}
