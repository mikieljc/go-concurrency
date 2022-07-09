package main

import "log"

func generateMessage(ch1 chan<- string) {
	//send
	ch1 <- "New Message"
}

func relayMessage(ch1 <-chan string, ch2 chan<- string) {
	//receive
	msg := <-ch1
	//send
	ch2 <- msg

	close(ch2)
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go generateMessage(ch1)
	go relayMessage(ch1, ch2)

	//receive and print

	msg := <-ch2

	log.Println(msg)
}
