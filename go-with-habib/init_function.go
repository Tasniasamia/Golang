package main

import "fmt"

var x = 10

func init() {
    fmt.Println("init")
}

func Add(a, b int) int {
    fmt.Println("Add")
    return a + b
}
func Sub(a, b int) int {
    fmt.Println("Sub")
    return a -b
}


func main() {
    Add(1, 2);
	Sub(2,1);
	fmt.Println("main")

}