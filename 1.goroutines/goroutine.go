package goroutines

import "fmt"

func Run() {
	go helloworld()
	// continue other code
	// ...

	fmt.Println("this run from Run() func")

	go func() {
		fmt.Println("this call goroutine from anonymous function")
	}()
}

func helloworld() {
	fmt.Println("this call goroutine from a function")
}
