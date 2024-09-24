package service

import (
	"fmt"
)

type PubSub interface {
	Publish()
	Subscribe()
}

type Event struct {
	EventID int
	Data    string
}

type PubSubHelper struct {
	ch chan Event
	// wg sync.WaitGroup
}

func CreatePubSub() PubSubHelper {
	helper := PubSubHelper{
		ch: make(chan Event),
	}
	return helper
}

func (helper PubSubHelper) Publish() {

	for i := 0; i < 10; i++ {
		event := Event{
			EventID: i,
			Data:    "Some data",
		}
		helper.ch <- event

	}
	close(helper.ch)
}

func (helper PubSubHelper) Subscribe() {

	for event := range helper.ch {
		fmt.Println("Received event:", event)
	}
}
