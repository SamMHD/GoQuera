package main

import (
	"fmt"
	"sync"
	"time"
)

func Solution(d time.Duration, message string, ch ...chan string) (numberOfAccesses int) {
	numberOfAccesses = 0
	var wg sync.WaitGroup
	wg.Add(len(ch))

	for i := range ch {
		fmt.Println("Creating go routine for channel", i)
		go func(i int) {
			defer wg.Done()
			fmt.Println("Sending message to channel", i)
			select {
			case <-time.After(d * time.Second):
				fmt.Println("Timeout", i)
				return
			case ch[i] <- message:
				fmt.Println("Message sent to channel", i)
				numberOfAccesses++
				return
			}
		}(i)
	}

	wg.Wait()
	return numberOfAccesses
}

// func main() {
// 	a := make(chan string, 1)
// 	b := make(chan string, 1)
// 	b <- "salam"

// 	go func() {
// 		<-time.After(3 * time.Second)
// 		s := <-b
// 		fmt.Println("k", s)
// 	}()

// 	res := Solution(1, "ali", a, b)
// 	fmt.Println(res)
// }
