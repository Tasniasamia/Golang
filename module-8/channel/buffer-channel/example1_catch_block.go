package main;

import (
	"fmt"
	"time"
	"sync"
)
var wg sync.WaitGroup;



func main(){
buffer_channel :=make(chan string,2);
wg.Add(1)
go func(){
defer wg.Done();
time.Sleep(2*time.Second);
buffer_channel <- "I am funcion 1";
}()

wg.Add(1)
go func(){
defer wg.Done();
time.Sleep(1*time.Second);
buffer_channel <- "I am function 2";
}()

wg.Add(1)
go func(){
defer wg.Done();
time.Sleep(3*time.Second);
buffer_channel <- "I am function 3";
}()

wg.Add(1)
go func(){
defer wg.Done();
time.Sleep(3*time.Second);
buffer_channel <- "I am function 4";
}()



for range 1{
fmt.Println(<- buffer_channel);
}
wg.Wait();
fmt.Println("Main End");


}