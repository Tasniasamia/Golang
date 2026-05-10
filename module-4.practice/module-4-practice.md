# Go Practice Problems 🐹

তুমি এই topics গুলো শিখেছো — সেগুলোর উপর ভিত্তি করে নিচের problems গুলো solve করো।

---

## 📦 Topic 1: Map

### Problem 1 — Student Marks
একটা map বানাও যেখানে student এর নাম key এবং marks value।
- ৩ জন student এর নাম ও marks দাও
- সবার marks print করো
- একজনের marks update করো
- একজনকে delete করো
- শেষে loop দিয়ে সব print করো

### Problem 2 — Word Counter
একটা string দেওয়া আছে: `"go is fun and go is easy"`
- প্রতিটা word কতবার আছে সেটা map দিয়ে count করো
- Output হবে: `go:2, is:2, fun:1, and:1, easy:1`

---

## 📦 Topic 2: Array & Slice

### Problem 3 — Array Sum
`[5]int{10, 20, 30, 40, 50}` এই array এর সব element এর sum বের করো।

### Problem 4 — Slice Modify
`[]string{"Tasnia", "Tisha", "Arin"}` এই slice এ loop দিয়ে প্রতিটা নামের শেষে `" Sharin"` যোগ করো।
- **Hint:** `range` এর value দিয়ে হবে না, index ব্যবহার করো!

### Problem 5 — Filter Slice
`[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}` এই slice থেকে শুধু জোড় সংখ্যাগুলো নিয়ে নতুন slice বানাও।

---

## 📦 Topic 3: String, Byte & Rune

### Problem 6 — Character Count
একটা string নাও: `"Tasnia Sharin"`
- loop দিয়ে প্রতিটা character ও তার ASCII value print করো
- Total কতটা character আছে সেটা print করো

### Problem 7 — Byte vs Rune
`"Hello"` string টাকে:
- `[]byte` এ convert করে print করো (type সহ)
- `range` দিয়ে loop করে রুন print করো (type সহ)
- পার্থক্য comment এ লিখো

### Problem 8 — Vowel Counter
একটা string এ কতটা vowel (a, e, i, o, u) আছে সেটা count করো।

---

## 📦 Topic 4: Struct & Method

### Problem 9 — Bank Account
একটা `bankAccount` struct বানাও যেখানে:
- `owner` string
- `balance` float64

Methods:
- `deposit(amount float64)` — balance বাড়াবে
- `withdraw(amount float64)` — balance কমাবে (balance 0 এর নিচে যেতে পারবে না)
- `getBalance()` — balance print করবে

### Problem 10 — Rectangle
একটা `rectangle` struct বানাও:
- `width`, `height` float64

Methods:
- `area()` float64 — ক্ষেত্রফল return করবে
- `perimeter()` float64 — পরিসীমা return করবে
- `isSquare()` bool — width == height হলে true

---

## 📦 Topic 5: Interface

### Problem 11 — Shape Interface
একটা `shape` interface বানাও:
```go
type shape interface {
    area() float64
    perimeter() float64
}
```
- `circle` struct বানাও (radius float64)
- `rectangle` struct বানাও (width, height float64)
- দুইটাই `shape` implement করো
- একটা function বানাও `printShapeInfo(s shape)` যেটা area ও perimeter print করবে

### Problem 12 — Payment System (bKash & Rocket)
তোমার শেখা bKash/Nagod pattern এর মতো:
- `payment` interface বানাও `pay(amount float64)` method সহ
- `bkash`, `rocket`, `nagod` struct বানাও
- প্রতিটার আলাদা `pay()` implement করো
- `checkout(p payment, amount float64)` function বানাও
- তিনটা payment method দিয়ে test করো

### Problem 13 — Struct এর ভেতরে Interface
একটা `person` struct বানাও:
- `name` string
- `vehicle` — একটা interface (যেটায় `drive()` method আছে)

`car` এবং `bike` struct বানাও, দুইটাই `vehicle` implement করো।
- person1 এর vehicle হবে car
- person2 এর vehicle হবে bike
- দুইজনকেই drive করাও

---

## 📦 Topic 6: Mixed Challenge

### Problem 14 — Student Database
`map[string]struct` ব্যবহার করে একটা student database বানাও:
- struct এ থাকবে: `name`, `age`, `grade`
- ৩ জন student add করো
- সবার info print করো (loop দিয়ে)
- একজনের grade update করো

### Problem 15 — Animal Shelter
- `animal` interface বানাও: `speak()`, `eat()` methods সহ
- `cat`, `dog`, `bird` struct বানাও
- একটা `[]animal` slice বানাও (shelter)
- সব animal কে loop দিয়ে speak ও eat করাও

---

## 💡 Tips
- Pointer receiver (`*struct`) ব্যবহার করো যখন struct modify করবে
- Interface এ pass করার সময় `&` দিতে হবে (যদি pointer receiver হয়)
- `range` এ value copy আসে, modify করতে index ব্যবহার করো
- Map এ `delete(map, key)` দিয়ে element মুছা যায়