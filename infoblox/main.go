package main

import "fmt"

// 1. Odd even with sequence like 1 3 2 5 7 4
// 2. What is CRD ? what exactly happens in the background ?
// 3. How to find duplicate in sequence of numbers in O(N) [1, 100]
// 4. Odd even with sequence like 1 3 2 5 7 4 (with single go routine

// even, odd
//
//	generator
//	buffered channel
//	two workers consuming
//	even - consume two values
//	odd - consume a single value
func main() {

	/*
		oddCh := make(chan int)
		evenCh := make(chan int)

		go generator(oddCh, evenCh)

		for oddCh != nil || evenCh != nil {
			select {
			case value, oddClosed := <-oddCh:
				if oddClosed {
					fmt.Println("Odd: ", value)
				} else {
					oddCh = nil
				}
			case value, evenClosed := <-evenCh:
				if evenClosed {
					fmt.Println("Even: ", value)
				} else {
					evenCh = nil

				}
			}
		}
	*/

	oddCh := make(chan int)
	evenCh := make(chan int)

	go generator1(oddCh, evenCh)

	for oddCh != nil || evenCh != nil {
		select {
		case value, oddClosed := <-oddCh:
			if oddClosed {
				fmt.Println("Odd: ", value)
			} else {
				oddCh = nil
			}
		case value, evenClosed := <-evenCh:
			if evenClosed {
				fmt.Println("Even: ", value)
			} else {
				evenCh = nil

			}
		}
	}
}

func generator1(oddCh, evenCh chan int) {
	var evenCounterValue []int
	oddCounter := 0
	for i := 1; i <= 10; i++ {
		if i%2 != 0 && oddCounter != 2 {
			oddCounter += 1
			oddCh <- i
		} else if i%2 != 0 && oddCounter == 2 {
			evenCh <- evenCounterValue[0]
			evenCounterValue = evenCounterValue[1:]
			oddCounter = 1
			oddCh <- i
		} else {
			evenCounterValue = append(evenCounterValue, i)
		}
	}
	for _, value := range evenCounterValue {
		evenCh <- value
	}
	close(evenCh)
	close(oddCh)
}

// func generator(oddCh, evenCh chan int) {

// 	evenOddSignal := make(chan struct{})
// 	go func() {
// 		for i := 2; i <= 10; i = i + 2 {
// 			_, isOpen := <-evenOddSignal
// 			if !isOpen {
// 				evenCh <- i
// 			} else {
// 				evenCh <- i
// 				evenOddSignal <- struct{}{}
// 			}
// 		}
// 		close(evenCh)
// 	}()

// 	go func() {
// 		oddLen := 0
// 		<-evenOddSignal
// 		for i := 1; i <= 10; i = i + 2 {
// 			if oddLen != 2 {
// 				oddCh <- i
// 				oddLen += 1
// 			} else {
// 				evenOddSignal <- struct{}{}
// 				oddLen = 1
// 				<-evenOddSignal
// 				oddCh <- i
// 			}
// 		}
// 		close(oddCh)
// 		close(evenOddSignal)
// 	}()

// 	evenOddSignal <- struct{}{}

// }
