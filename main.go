package main

import (

	"fmt"
	"math"
	"math/rand"
)

//add new function

func addition (a int, b int) int  {
	
	return a + b
}


func main() {

	println("Hello word")

	println("My fovotite Number Is :", rand.Intn(24))

	println("Now you Have %g Problems ", math.Sqrt(91))

	fmt.Println(addition(45,67))
	fmt.Printf("Now you Have %g Problems ", math.Sqrt(91))
}