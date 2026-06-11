package main;
import "fmt";

func main(){
slice :=make([]int,3,5);

slice[0]=1;
slice[1]=11;
slice[2]=12;
slice=append(slice,13);
slice[3]=132;
// slice[22]=22;  not acceptable
slice=append(slice,14);

slice=append(slice,15);

fmt.Println(slice);
}






// চলুন এই code-এর simulation সম্পর্কে জানা যাক।

// Step 1:

// main() function call হওয়ার সাথে সাথে stack এ main frame create হবে।

// Step 2:
// slice := make([]int,3,5)

// এর জন্য stack-এর ভিতরে একটি slice header তৈরি হবে, যেখানে থাকবে:

// pointer (underlying array-এর address)
// length = 3
// capacity = 5

// Compiler escape analysis করে দেখবে যে make() দিয়ে একটি নতুন slice তৈরি করা হচ্ছে এবং এর জন্য underlying array দরকার। তাই runtime heap-এ capacity 5 বিশিষ্ট একটি underlying array-এর জন্য memory allocate করবে।

// যেহেতু length = 3 এবং int-এর zero value হলো 0, তাই প্রথম 3টি position এ 0 থাকবে।

// Heap-এর অবস্থা:

// index:  [0] [1] [2] [3] [4]
// value:  [0] [0] [0] [0] [0]

// length = 3
// capacity = 5

// (বাস্তবে সবগুলো cell zero initialize হয়।)

// Stack-এর slice header pointer দিয়ে এই underlying array-কে point করবে।

// Step 3:
// slice[0] = 1

// Underlying array-এর index 0 তে 1 store হবে।

// [1,0,0,0,0]
// Step 4:
// slice[1] = 11
// slice[2] = 12

// একইভাবে value update হবে।

// [1,11,12,0,0]
// Step 5:
// slice = append(slice,13)

// বর্তমানে:

// length = 3
// capacity = 5

// অর্থাৎ index 3 এখনো খালি আছে।

// তাই append() call হলে নতুন array create করার দরকার হবে না।

// Underlying array-এর index 3 এ 13 store হবে।

// [1,11,12,13,0]

// এবং

// length = 4
// capacity = 5

// হবে।

// এরপর:

// slice[3] = 132

// index 3-এর value overwrite হয়ে যাবে।

// [1,11,12,132,0]
// Step 6:
// slice = append(slice,14)

// এখন:

// length = 4
// capacity = 5

// index 4 খালি আছে।

// তাই নতুন array লাগবে না।

// [1,11,12,132,14]

// এবং

// length = 5
// capacity = 5

// হবে।

// Step 7:
// slice = append(slice,15)

// এখন:

// length = 5
// capacity = 5

// Length এবং capacity সমান হয়ে গেছে।

// অর্থাৎ বর্তমান underlying array-এ আর কোনো খালি জায়গা নেই।

// তাই runtime নতুন একটি বড় underlying array allocate করবে heap-এ।

// Go runtime সাধারণত capacity বাড়িয়ে দেয় (এই ক্ষেত্রে সাধারণত 10 হবে)।

// নতুন array:

// [1,11,12,132,14,0,0,0,0,0]

// পুরনো array থেকে data copy হবে:

// [1,11,12,132,14,0,0,0,0,0]

// এরপর 15 append হবে:

// [1,11,12,132,14,15,0,0,0,0]

// এখন slice header update হবে:

// length = 6
// capacity = 10

// এবং pointer নতুন underlying array-কে point করবে।

// এই সিদ্ধান্ত Go runtime নেয়।

// পুরনো array আর কোনো reference না থাকলে পরে Garbage Collector সেটি remove করতে পারবে।

// Step 8:
// fmt.Println(slice)

// fmt.Println() frame create হবে।

// Slice print করার সময় Go slice-এর pointer থেকে length পর্যন্ত data print করবে।

// বর্তমানে:

// length = 6

// তাই print হবে:

// [1 11 12 132 14 15]

// এরপর fmt.Println() frame pop হবে।

// গুরুত্বপূর্ণ কথা
// slice = append(slice, value)

// এভাবে নতুন value add করা হয়।

// আর existing value update করতে হয়:

// slice[index] = newValue

// যেমন:

// slice[3] = 132

// কিন্তু

// slice[22] = 22

// দেওয়া যাবে না, কারণ slice-এর length 5 বা 6 হলে index 22 valid নয়। Program panic করবে:

// panic: runtime error: index out of range

// Final Output:

// [1 11 12 132 14 15]