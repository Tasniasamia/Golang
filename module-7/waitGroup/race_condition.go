package main;

import (
	"fmt"
	"sync"
)

var counter int;
var wg sync.WaitGroup;
var mu sync.Mutex
func main(){

	
	for i:=1;i<1000;i++{
	
        wg.Add(1);
		go increment();
	}
		

	wg.Wait();
	fmt.Println("Counter is ",counter);
}

func increment(){
	defer wg.Done()
    mu.Lock()
	defer mu.Unlock()
   counter=counter+1;
}