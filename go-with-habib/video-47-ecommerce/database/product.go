package database;

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Roll int    `json:"roll"`
	Age  int    `json:"age"`
	Dept string `json:"dept"`
}

var Users [] User;