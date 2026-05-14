package main;

import (
	"fmt"
	"time"
	"sync"
);


var wg sync.WaitGroup;

func uploadFile(c chan string){
defer wg.Done();
fmt.Println("Start Uploading....");
time.Sleep(3*time.Second);
fmt.Println("Uploaded and return URL");
c <- "https://example.org/image1.png"; 
}


func main(){
start :=time.Now();
channel1 :=make(chan string);
wg.Add(1);
go uploadFile(channel1);
wg.Wait();
// fmt.Println(<- channel1);
fmt.Println("Total time is ",time.Since(start));
}

// output
// Start Uploading....
// Uploaded and return URL
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [semacquire]:
// sync.runtime_Semacquire(0xc000096040?)
//         /usr/lib/go-1.22/src/runtime/sema.go:62 +0x25
// sync.(*WaitGroup).Wait(0x100530520?)
//         /usr/lib/go-1.22/src/sync/waitgroup.go:116 +0x48
// main.main()
//         /mnt/c/Users/HP/Desktop/Programming_Hero/Go/module-8/channel/create.channel.main_thread_not_wait.go:26 +0x92

// goroutine 20 [chan send]:
// main.uploadFile(0xc00009e060)
//         /mnt/c/Users/HP/Desktop/Programming_Hero/Go/module-8/channel/create.channel.main_thread_not_wait.go:17 +0xe7
// created by main.main in goroutine 1
//         /mnt/c/Users/HP/Desktop/Programming_Hero/Go/module-8/channel/create.channel.main_thread_not_wait.go:25 +0x86
// exit status 2