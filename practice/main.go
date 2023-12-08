package main

import (
	//"bufio"
	"fmt"
	//"math"
	/* 	"os"
	   	"strconv"
	   	"strings" */
	"time"
)

const num int = 54

func main() {
	/* var aString = "this is go"
	fmt.Println(aString)


	fmt.Printf("the variablels type is %T\n", aString)

	myString := "this is also a string\n"

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("you entered: ", input)

	//converting strings to other inputs
	fmt.Print("enter a number: ")
	numInput, _ := reader.ReadString('\n')

	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number: ", aFloat)
	}


		GO does not implicitely convert number types, so we have to convert if we are adding
		ints and floats
		We need to convert before using
		To do this, we wrap in a function of the same name of the target type


	var anInt int = 4
	var aFloatAgain float64 = 42
	sum := float64(anInt) + aFloatAgain

	fmt.Printf("Sum: %v, Type %T", sum, sum)


		For more complicated math, we have to go to the "math" package
	*/

	/* 	var aFloatThree = 3.14159
	   	var rounder = math.Round(aFloatThree)
	   	fmt.Printf("Original: %v, Rounded: %v\n", aFloatThree, rounder)

	   	i1, i2, i3 := 12, 45, 68

	   	intSum := i1 + i2 + i3

	   	fmt.Println(intSum)

	   	f1, f2, f3 := 23.5, 12.3, 45.6

	   	floatSum := f1 + f2 + f3

	   	fmt.Println(floatSum) //the result of this is imprecise because of fp arithmetic

	   	floatSum = math.Round(floatSum*100) / 100
	   	fmt.Println(floatSum) //doing it this way is the safest way to accurately calculate floats

	   	circleRadius := 15.5
	   	circumference := circleRadius * 2 * math.Pi
	   	fmt.Printf("Circumference: %.2f", circumference) //prints to the nearest 2nd digit */
	n := time.Now()
	fmt.Println(n)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	fmt.Printf("back it up,type is %T\n", parsedTime)

}
