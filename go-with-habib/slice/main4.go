package main;
import "fmt";

func main(){
	var arr[]int;
	arr=append(arr,1);
	arr=append(arr,2);
	arr=append(arr,4,5,6,7,8);
	fmt.Println(arr);

}


// Explanation:

// Step-1: Code segment এ থাকবে main() আর fmt.Println এর instructions।
// Step-2: Stack এ main frame create হবে।
// Step-3: var arr[]int — stack এ arr slice header create হবে। pointer=nil, length=0, capacity=0।
// Step-4: arr=append(arr,1)

// Stack এ append frame create হবে
// capacity=0, length=0 → জায়গা নেই, heap এ capacity=1 এর নতুন underlying array create হবে → {1}
// length=1, capacity=1, pointer → নতুন array এর address
// append frame pop হবে

// Step-5: arr=append(arr,2)

// Stack এ append frame create হবে
// capacity=1, length=1 → জায়গা নেই, heap এ capacity=2 এর নতুন underlying array create হবে
// পুরানো {1} copy হবে, তারপর 2 store হবে → {1,2}
// length=2, capacity=2, pointer change হবে
// পুরানো array GC করবে
// append frame pop হবে

// Step-6: arr=append(arr,4,5,6,7,8)

// Stack এ append frame create হবে
// capacity=2, length=2 → দরকার 2+5=7 জায়গা, capacity=2 → দরকার নেই 2→4→8, তাই heap এ capacity=8 এর নতুন underlying array create হবে
// পুরানো {1,2} copy হবে, তারপর {4,5,6,7,8} store হবে → {1,2,4,5,6,7,8}
// length=7, capacity=8, pointer change হবে
// পুরানো array GC করবে
// append frame pop হবে

// Step-7: fmt.Println(arr)

// Stack এ fmt.Println frame create হবে
// slice print হবে
// frame pop হবে