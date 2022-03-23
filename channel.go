package main

import (
	"fmt"
	"time"
)

func main() {
	// c := make(chan string)

	// c <- value

	// result := <- c

	// c := make(chan string)

	// go introduce("Raditeo", c)
	// go introduce("Warma", c)

	// msg1 := <-c
	// fmt.Println(msg1)

	// msg2 := <-c
	// fmt.Println(msg2)

	// close(c)

	// c := make(chan string)

	// students := []string{"Raditeo", "Warma"}

	// for _, v := range students {
	// 	go func(student string) {
	// 		fmt.Println("Student", student)
	// 		result := fmt.Sprintf("hi my name is %s", student)
	// 		c <- result
	// 	}(v)
	// }

	// for i := 1; i <= 2; i++ {
	// 	print(c)
	// }

	// close(c)

	// c := make(chan string)

	// students := []string{"Raditeo", "Warma"}

	// for _, v := range students {
	// 	go introduce(v, c)
	// }

	// for i := 1; i <= 2; i++ {
	// 	print(c)
	// }

	// close(c)

	// c1 := make(chan int)

	// go func(c chan int) {
	// 	fmt.Println("func goroutine starts sending data into the channel")
	// 	c <- 10
	// 	fmt.Println("func goroutine after sending data into the channel")
	// }(c1)

	// fmt.Println("main goroutine sleeps for 2 seconds")
	// time.Sleep(time.Second * 2)

	// fmt.Println("main goroutine starts receving data")
	// d := <-c1
	// fmt.Println("main goroutine received data:", d)

	// close(c1)
	// time.Sleep(time.Second)

	// c1 := make(chan int, 5)

	// go func(c chan int) {
	// 	for i := 1; i <= 5; i++ {
	// 		fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
	// 		c <- i
	// 		fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
	// 	}
	// 	close(c)
	// }(c1)

	// fmt.Println("main goroutine sleeps for 2 seconds")
	// time.Sleep(time.Second * 2)

	// for v := range c1 {
	// 	fmt.Println("main goroutine received value from channel:", v)
	// }

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		c1 <- "Hello!"
	}()

	go func() {
		time.Sleep(1 * time.Second)

		c2 <- "Salute!"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

// func introduce(student string, c chan string) {
// 	result := fmt.Sprintf("hi, my name is %s", student)

// 	c <- result
// }

// func print(c chan string) {
// 	fmt.Println(<-c)
// }

// func introduce(student string, c chan<- string) {
// 	result := fmt.Sprintf("hi, my name is %s", student)

// 	c <- result
// }

// func print(c <-chan string) {
// 	fmt.Println(<-c)
// }
