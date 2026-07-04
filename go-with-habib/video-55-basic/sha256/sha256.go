package main

import (
	"crypto/sha256"
	"fmt"
)

func main(){
    byteArr := []byte("Hello World")
	sha256Hash := sha256.Sum256(byteArr);
	fmt.Println(sha256Hash)
}