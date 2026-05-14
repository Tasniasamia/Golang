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

// output:
// I am function 2
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [semacquire]:
// sync.runtime_Semacquire(0x10?)
//         /usr/lib/go-1.22/src/runtime/sema.go:62 +0x25
// sync.(*WaitGroup).Wait(0x4bd8e8?)
//         /usr/lib/go-1.22/src/sync/waitgroup.go:116 +0x48
// main.main()
//         /mnt/c/Users/HP/Desktop/Programming_Hero/Go/module-8/channel/buffer-channel/example1.deadlock-1.go:47 +0x1f1

// goroutine 9 [chan send]:
// main.main.func4()
//         /mnt/c/Users/HP/Desktop/Programming_Hero/Go/module-8/channel/buffer-channel/example1.deadlock-1.go:39 +0x6f
// created by main.main in goroutine 1
//         /mnt/c/Users/HP/Desktop/Programming_Hero/Go/module-8/channel/buffer-channel/example1.deadlock-1.go:36 +0x165
// exit status 2