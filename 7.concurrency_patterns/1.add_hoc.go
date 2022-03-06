package concurrencypatterns

import "fmt"

func AddHocSample() {
	data := make([]int, 4) // this must be untouchable in convention
	// by convention we’re only accessing data from the loopData function
	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}
	handleData := make(chan int)
	go loopData(handleData)
	for num := range handleData {
		fmt.Println(num)
	}
	// But as the code is touched by many people, and deadlines loom,
	// mistakes might be made, and the confinement might break down and cause issues
}
