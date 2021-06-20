package main

import (
	"fmt"
	"sort"
	
)


func main() {
	// Stings
	// var nameOne string = "Mario"
	// var nameTwo = "Luigi"
	// var nameThree string

	// nameOne = "peach"
	// nameThree = "Bowser"
	
	// fmt.Println(nameOne, nameTwo)

	// nameFour := "Yoshi"
	// fmt.Println(nameFour)

	// Ints
	// var ageOne int = 20
	// var ageTwo = 30
	// ageThree := 40

	// fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory 
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint = 256

	// var scoreOne float32 = 25.98
	// var scoreTwo float64 = 3247329847328947.9
	// scoreThree := 1.5

	// Print
	// age := 35
	// name := "Shaun"

	// fmt.Print("Hello, ")
	// fmt.Print("world! \n")
	// fmt.Print("new line \n")

	// fmt.Println("hello!")
	// fmt.Println("goodbye!")
	// fmt.Println("My age is", age, "and my name is", name)

	// Printf (formatted strings) %_ = format specifier
	// fmt.Printf("My age is %v and my name is %v \n", age, name)
	// fmt.Printf("My age is %q and my name is %q \n", age, name)
	// fmt.Printf("age is of type %T", age)
	// fmt.Printf("you scored %f points! \n", 225.55)
	// fmt.Printf("you scored %0.1f points! \n", 225.55)

	// Sprintf (save formatted strings)
	// var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)
	// fmt.Println("The saved string is:", str)

	// Arrays and slices
	// var ages [3]int = [3]int{20, 25, 30}
	// var ages = [3]int{20, 25, 30}

	// names := [4]string{"Yoshi", "Mario", "Peach", "Bowser"}
	// names[1] = "Luigi"

	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	// Slices (use arrays under the hood)
	// var scores = []int{100, 50, 60}
	// scores[2] = 25
	// scores = append(scores, 85)

	// fmt.Println(scores, len(scores))

	// Slice Ranges
	// rangeOne := names[1:3]
	// rangeTwo := names[2:]
	// rangeThree := names[:3]

	// fmt.Println(rangeOne, rangeTwo, rangeThree)

	// rangeOne = append(rangeOne, "koopa")
	// fmt.Println(rangeOne)

	// Packages
	// greeting := "hello there friends!"

	// fmt.Println(strings.Contains(greeting, "hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "ll"))
	// fmt.Println(strings.Split(greeting, " "))

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser"))
}