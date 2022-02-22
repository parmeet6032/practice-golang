package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.stackoverflow.com",
		"https://go.dev/",
		"https://www.amazon.com",
	}

	// // one by one
	// for _, link := range links {
	// 	checkLink(link)
	// }

	// // to implement concurrency
	// c := make(chan string)

	// for _, link := range links {
	// 	go checkLink(link, c)
	// }

	// // fmt.Println(<-c) // waits for a message from channel

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// to continuosly check
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for {	// infinite loop
	// 	go checkLink(<-c, c)
	// }

	// for l := range c { // alternative loop, detects each msg from channel
	// 	time.Sleep(5 * time.Second)
	// 	go checkLink(l, c)
	// }

	for l := range c { // alternative loop, detects each msg from channel
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// continuos checking
func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

// // to implement concurrency
// func checkLink(link string, c chan string) {
// 	_, err := http.Get(link)

// 	if err != nil {
// 		fmt.Println(link, "might be down!")
// 		fmt.Println("* * * * * * * * * ERROR * * * * * * * * * \n", err, "\n* * * * * * * * * * * * * * * * * * * * * ")
// 		c <- "Might be down I think!"
// 		return
// 	}

// 	fmt.Println(link, "is up!")
// 	c <- "Yep its up!"
// }

// // to check one by one
// func checkLink(link string) {
// 	_, err := http.Get(link)

// 	if err != nil {
// 		fmt.Println(link, "might be down!")
// 		fmt.Println("********* ERROR:- \n", err, "\n***********************ß")
// 		return
// 	}

// 	fmt.Println(link, "is up!")
// }
