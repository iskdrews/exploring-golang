package main

import "fmt"

func main()  {
	fmt.Println("Hello everyone")

	foo()

	fmt.Println("Something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("Im in foo")
}

func bar()  {
	fmt.Println("And then we exited")
}