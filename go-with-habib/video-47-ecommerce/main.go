package main

import (
"mains/cmd"
"mains/database"
)










func main() {
	cmd.Start()
}

func init() {
	u1 := database.User{
		Id:   1,
		Name: "John Doe",
		Roll: 101,
		Age:  25,
		Dept: "Computer Science",
	}
	u2 := database.User{
		Id:   2,
		Name: "Jane Smith",
		Roll: 102,
		Age:  23,

		Dept: "Mathematics",
	}
	u3 := database.User{
		Id:   3,
		Name: "Bob Johnson",
		Roll: 103,
		Age:  24,
		Dept: "Physics",
	}
	database.Users = append(database.Users, u1, u2, u3)
}