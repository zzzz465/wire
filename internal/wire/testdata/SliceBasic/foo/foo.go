package main

import "fmt"

func main() {
	msgs := initMessages()
	for _, m := range msgs {
		fmt.Println(m)
	}
}

type Messages []string

func provideHello() string { return "hello" }
func provideWorld() string { return "world" }
