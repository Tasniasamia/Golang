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



for range 5{
fmt.Println(<- buffer_channel);
}
wg.Wait();
fmt.Println("Main End");


}

// output:
// I am function 2
// I am funcion 1
// I am function 4
// I am function 3
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan receive]:
// main.main()
//         /mnt/c/Users/HP/Desktop/Programming_Hero/Go/module-8/channel/buffer-channel/example1.deadlock-2.go:45 +0x185
// exit status 2