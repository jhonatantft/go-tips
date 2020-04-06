package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// a goroutine é independente e desconhecida pela thread principal
	go say("world")
	say("Hello") // Sem goroutine
}