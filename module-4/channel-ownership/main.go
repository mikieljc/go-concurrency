package main

import (
	"fmt"
	"strconv"
)

// Best Practices in using channels
//1. The owner of channel should be a "GO ROUTINE" that instantiates, writes & closes a channel
//2. The utilizers of channel should only have a "READ ONLY" view into the channel

//This tips will help us avoid the following:
//1. Deadlocking by writing to a nil channel
//2. closing a channel (causes panic)
//3. writing to a closed channel (causes panic)
//4. closing a channel more than once (causes panic)

type Channel interface {
	produceMessage()
	getChannel() <-chan string
}

type channel struct {
	ch chan string
}

func (c *channel) produceMessage() {
	defer close(c.ch)
	for i := 0; i < 10; i++ {
		c.ch <- "Message #" + strconv.Itoa(i)
	}
}

func (c *channel) getChannel() <-chan string {
	return c.ch
}

func NewChannelOwner() Channel {
	ch := make(chan string)

	return &channel{
		ch: ch,
	}
}

func main() {

	chOwner := NewChannelOwner()

	go chOwner.produceMessage()

	rcvChannel := chOwner.getChannel()

	consumer(rcvChannel)

}

func consumer(ch <-chan string) {
	for msg := range ch {
		fmt.Println("RECEIVED MESSAGE: ", msg)
	}
	fmt.Println("DONE RECEIVING!")
}
