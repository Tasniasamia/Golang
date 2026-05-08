package main;

import "fmt";

func main(){

	studentMarks:=[10] int {};
   pass:=0;
   fail:=0;
   max:=0;
	for i:=0;i<10;i++{
		fmt.Scan(&studentMarks[i]);

		if(studentMarks[i]>max){
			max=studentMarks[i];
		}

		if(studentMarks[i]>=33){
        pass++;
		}else{
		fail++;
        }
	}
	fmt.Println("Total fail ",fail);
	fmt.Println("Total Pass ",pass);
	fmt.Println("Highest Number is ",max);


}