package main;
import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main(){
key := []byte("mysecret")
message := []byte("Hello")

h := hmac.New(sha256.New, key)
h.Write(message);
fmt.Println(h.Sum(nil));
}