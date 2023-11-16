package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sm(ch chan int, nf int) {
	seconds := rand.Intn(10)
	ss := time.Duration(seconds)
	time.Sleep(time.Second * ss)
	fmt.Printf("номер вызова - %v время %v\n", nf, ss)
	ch <- int(ss)
}

func main1() {
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go sm(ch, i)
	}

	for i := 0; i < 5; i++ {
		someEl := <-ch
		fmt.Printf("someEl: %v\n", someEl)
	}

}
