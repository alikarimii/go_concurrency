package concurrencypatterns

func Generator(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for _, i := range integers {
			select {
			case <-done:
				return
			case intStream <- i:
			}
		}
	}()
	return intStream
}
func Sq(done <-chan interface{}, in <-chan int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := range in {
			select {
			case <-done:
				return
			case intStream <- i * i:
			}
		}
	}()
	return intStream
}

func Multiply(
	done <-chan interface{}, intStream <-chan int, multiplier int,
) <-chan int {
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for i := range intStream {
			select {
			case <-done:
				return
			case multipliedStream <- i * multiplier:
			}
		}
	}()
	return multipliedStream
}

func Add(
	done <-chan interface{}, intStream <-chan int, additive int,
) <-chan int {
	addedStream := make(chan int)
	go func() {
		defer close(addedStream)
		for i := range intStream {
			select {
			case <-done:
				return
			case addedStream <- i + additive:
			}
		}
	}()
	return addedStream
}
