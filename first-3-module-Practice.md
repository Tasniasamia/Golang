# Golang Practice Problems
### Module 63, 64 & 65

---

## Module 63 — Basics

### Practice 1: Variables & Constants
1. তোমার নাম, বয়স, GPA — এই ৩টা variable declare করো
2. PI = 3.1416 constant বানাও
3. সব Printf দিয়ে print করো

**Expected Output:**
```
Name: Samia
Age: 20
GPA: 3.85
PI: 3.1416
```

---

### Practice 2: Data Types & Zero Values
1. int, float64, string, bool — এই ৪ type এর variable declare করো (value দিও না)
2. Zero value কী আসে print করো
3. Printf তে %T দিয়ে প্রতিটার type print করো

**Expected Output:**
```
int zero value: 0  (type: int)
float64 zero value: 0  (type: float64)
string zero value:   (type: string)
bool zero value: false  (type: bool)
```

---

### Practice 3: Functions
1. দুইটা number নিয়ে sum return করে এমন function বানাও
2. একটা number নিয়ে সেটা Even না Odd return করে এমন function বানাও
3. Named return দিয়ে rectangle এর area আর perimeter একসাথে return করো

**Hint — Named Return:**
```go
func rectInfo(length, width float64) (area float64, perimeter float64) {
    area = length * width
    perimeter = 2 * (length + width)
    return
}
```

---

## Module 64 — Control Flow

### Practice 4: If-Else
User থেকে বয়স নাও এবং নিচের শর্ত অনুযায়ী print করো:
- 18 এর কম হলে → "Minor"
- 18 থেকে 60 হলে → "Adult"
- 60 এর বেশি হলে → "Senior"

---

### Practice 5: Switch

**Problem A:**
1. User থেকে weekday number নাও (1-7)
2. Switch দিয়ে day এর নাম print করো (1=Monday, 2=Tuesday ...)

**Problem B — Normal Switch:**
```
score >= 80  →  "A"
score >= 60  →  "B"
score >= 40  →  "C"
বাকি সব     →  "Fail"
```

---

### Practice 6: For Loop
1. 1 থেকে 100 পর্যন্ত সব সংখ্যার যোগফল বের করো
2. User থেকে number নাও, এর multiplication table print করো (1-10)
3. 1 থেকে 50 এর মধ্যে শুধু odd numbers print করো (`continue` ব্যবহার করতে হবে)

---

### Practice 7: Problem Solving (fmt.Scan)
User থেকে ৫টা number নাও এবং:
- সর্বোচ্চ সংখ্যা বের করো
- সর্বনিম্ন সংখ্যা বের করো
- সব সংখ্যার গড় বের করো

> ⚠️ মনে রাখো: fmt.Scan এ space দেওয়া যাবে না!

---

## Module 65 — Arrays, Slices, Pointers & Structs

### Practice 8: Arrays
1. ১০টা student এর marks নাও array তে
2. Pass (>= 33) আর Fail এর count করো
3. সবচেয়ে বেশি marks কার বের করো

**Expected Output:**
```
Total Pass: 7
Total Fail: 3
Highest Marks: 95
```

---

### Practice 9: Slices
1. একটা empty slice বানাও
2. append() দিয়ে ৫টা number যোগ করো
3. Slice এর length আর capacity print করো
4. Slice এর মধ্যে থেকে index 1 থেকে 3 আলাদা করে print করো

**Hint:**
```go
nums := []int{}
nums = append(nums, 10)
subSlice := nums[1:4]
```

---

### Practice 10: Pointers

**Problem A:**
1. একটা variable এর value function দিয়ে change করো — প্রথমে pointer ছাড়া
2. তারপর pointer দিয়ে করো এবং পার্থক্য দেখো

**Problem B:**
1. Array pointer function এ pass করো
2. Function এর ভেতর থেকে সব value দ্বিগুণ করো
3. Main এ print করলে দ্বিগুণ হয়েছে কিনা দেখো

---

### Practice 11: Structs ⭐ (সবচেয়ে গুরুত্বপূর্ণ)
একটা Bank Account System বানাও।

**Struct:**
```go
type bankAccount struct {
    accountHolder string
    accountNumber int
    balance       float64
}
```

**করতে হবে:**
1. Constructor function দিয়ে account খোলো (`createAccount`)
2. Deposit function বানাও
3. Withdraw function বানাও (balance কম থাকলে error দেখাবে)
4. Balance check করার function বানাও

> ⚠️ মনে রাখো: struct, constructor এবং সব function অবশ্যই main এর **বাইরে** থাকবে।

**Expected Output:**
```
Account created for: Samia (No: 1001)
Deposit: 5000 → New Balance: 5000.00
Withdraw: 2000 → New Balance: 3000.00
Error: Insufficient balance!
Current Balance: 3000.00
```

---

## 🏆 Final Project — Student Management System
সব কিছু একসাথে ব্যবহার করো।

**Struct:**
```go
type Student struct {
    name  string
    roll  int
    age   int
    marks [3]float64  // ৩টা subject
}
```

**করতে হবে:**
1. ৫ জন student এর data নাও (fmt.Scan দিয়ে)
2. Constructor function দিয়ে student বানাও
3. প্রতিটার average marks বের করো
4. সবার মধ্যে highest scorer কে বের করো
5. সব student এর info সুন্দরভাবে print করো

**Expected Output:**
```
--- Student List ---
1: Samia  | Roll: 101 | Age: 20 | Avg: 85.33
2: Rafi   | Roll: 102 | Age: 21 | Avg: 78.00
...

Highest Scorer: Samia (85.33)
```

---

## 💡 Tips
- প্রতিটা practice নিজে করার চেষ্টা করো আগে
- fmt.Scan এ space যুক্ত input দেওয়া যাবে না
- struct, constructor এবং সব function main এর বাইরে রাখো
- Go তে semicolon লাগে না
- আটকে গেলে Claude কে জিজ্ঞেস করো