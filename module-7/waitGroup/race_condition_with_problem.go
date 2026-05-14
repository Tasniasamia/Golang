package main;

import (
	"fmt"
	"sync"
)

var counter int;
var wg sync.WaitGroup;

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

   counter=counter+1;
}