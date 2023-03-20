package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) {

	for {
		time.Sleep(delay)
		fmt.Printf("The time is %02d.%02d.%02d: %v\n", time.Now().Hour(), time.Now().Minute(), time.Now().Second(), text)
	}

}

func main() {
	go Remind("Time to eat", 10*time.Second)
	go Remind("Time to work", 30*time.Second)
	go Remind("Time to sleep", 60*time.Second)
	select {}
}
