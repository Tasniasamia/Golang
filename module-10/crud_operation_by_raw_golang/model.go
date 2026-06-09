package main;

type UserType struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cgpa float64 `json:"cgpa"`
	Age  int     `json:"age"`
}

var users = []UserType{
	{
		Id:   1,
		Name: "samia",
		Cgpa: 3.88,
		Age:  23,
	}, {
		Id:   2,
		Name: "ayla",
		Cgpa: 3.22,
		Age:  23,
	},
}
