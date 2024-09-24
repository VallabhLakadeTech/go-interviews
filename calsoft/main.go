package main

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

type APICall struct {
	result_count int
	results      []interface{}
}

func main() {

	details := make(map[string]APICall)

	gson.Get("action")
}
