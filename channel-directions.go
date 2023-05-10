package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
	fmt.Println("Sent to pings channel")
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	fmt.Println("Received from pings channel")
	pongs <- msg
	fmt.Println("Sent to Pongs channel")
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	go ping(pings, "Message stored")
	go pong(pings, pongs)
	fmt.Println("Original message: ", <-pongs)
	fmt.Println("Received from pongs channel")

}
