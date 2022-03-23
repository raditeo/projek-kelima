package main

import (
	"fmt"
	"os"
)

func main() {
	// defer fmt.Println("defer func starts to execute")
	// fmt.Println("Hi everyone")
	// fmt.Println("Welcome back to Go learning center")

	// callDeferFunc(true)
	// fmt.Println("Hi everyone!")

	defer fmt.Println("invoke with defer")
	fmt.Println("before exiting")
	os.Exit(1)
}

// func callDeferFunc(condition bool) {

// 	if !condition {
// 		fmt.Println("dont call deferFunc")
// 		return
// 	}

// 	defer deferFunc()
// }

// func deferFunc() {
// 	fmt.Println("defer func starts to execute")
// }
