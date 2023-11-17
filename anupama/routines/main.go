package main

import (
	"fmt"
	"net/http"
	"time"
)

// go routines are always used with Functions

// even though we are launching multiple go routines,
//only one is being executed or running at any given time.
// So essentially, even though we are spawning multiple go routines,
// they are not actually being executed truly at the same time, whenever we have one CPU.
// So this one CPU is only running the code inside of one go routine at a time.
// And we rely upon this go scheduler to decide which go routine is being executed.
// When we have multiple CPU cores, each one can run one single go routine at a time.
// And so the go scheduler might say, we've got three separate go routines and
// we have three separate CPU cores.
// So rather than monitoring each go routine and attempting to run only one at a time,
// the scheduler will instead assign one routine to this core,
// another one, the second core, and the last one to the third.
// we have multiple CPU cores, hen we're talking about running multiple chunks of code truly at the same time.

// CONCURRENCY : (kind of switching) So when we say something is concurrent,
// we are simply saying that our program has the ability to run
// different things kind of at the same time, but not really at the same time.
// Because when we have one core, we are only picking one go routine.

// PARALLELISM : We are literally saying that
// we can do multiple things at the exact same, like nanosecond.

func main() {
	links := []string{"http://google.com", "http://facebook.com",
		"http://stackoverflow.com", "http://golang.org", "http://amazon.com"}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		time.Sleep(5 * time.Second)
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // passing adress
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is ok")
	c <- link
}

// var wg sync.WaitGroup
// wg.Add(1)
// go printIntegers(&wg)
// wg.Add(1)
// go printCharacters(&wg)

// wg.Wait()

// func printIntegers(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 20; i++ {
// 		fmt.Println(i)
// 	}
// }

// func printCharacters(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	str := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

// 	for _, v := range str {
// 		fmt.Println(v)
// 	}
// }
