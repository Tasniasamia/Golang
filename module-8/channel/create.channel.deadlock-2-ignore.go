package main

import (
	"fmt"
	"time"
)

func uploadFile(c chan string) {
	fmt.Println("Start Uploading...")

	time.Sleep(2 * time.Second)

	fmt.Println("Trying to send value...")

	c <- "https://example.com/file.png" // ❗ BLOCK HERE
	fmt.Println("This will NOT print immediately")
}

func main() {
	ch := make(chan string)

	go uploadFile(ch)

	time.Sleep(5 * time.Second)

	fmt.Println("Main finished")
}