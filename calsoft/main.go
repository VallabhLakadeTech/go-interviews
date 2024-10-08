package main

import (
	"fmt"
	"sync"
)

// func main() {

// 	pubsub := service.CreatePubSub()

// 	go pubsub.Publish()

// 	pubsub.Subscribe()

// }

// type CustomValue[T int | float64 | string] struct {
// 	version int
// 	value   T
// }

// type CustomMap[T int | float64 | string] struct {
// 	key         string
// 	customValue []CustomValue[T]
// }

// map := []CustomMap

// /*

// key: []{version:"",value:""}

// */

// type APICall struct {
// 	result_count int
// 	results      []interface{}
// }

// func main() {

// 	details := make(map[string]APICall)

// 	gson.Get("action")
// }

func main() {

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go printOdd(ch, &wg, 15)
	go printEven(ch, &wg, 15)
	wg.Wait()
}

func printOdd(ch chan int, wg *sync.WaitGroup, num int) {
	for i := 1; i <= num; i = i + 2 {
		<-ch
		fmt.Println("Odd: ", i)
		if i < num {
			ch <- 1

			continue
		}
		wg.Done()
	}

}

func printEven(ch chan int, wg *sync.WaitGroup, num int) {
	ch <- 1
	for i := 2; i <= num; i = i + 2 {
		<-ch
		fmt.Println("Even: ", i)
		if i < num {
			ch <- 1

			continue
		}
		wg.Done()
	}

}
