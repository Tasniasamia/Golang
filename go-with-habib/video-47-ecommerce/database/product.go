package database;

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Roll int    `json:"roll"`
	Age  int    `json:"age"`
	Dept string `json:"dept"`
}

var Users [] User;

func List() []User {
return Users;
}

func Store(user User) User {
	user.Id=len(Users)+1;
	Users = append(Users, user)
	return user
}

func Find(id int) User {
	for _, user := range Users {
		if user.Id == id {
			return user
		}
	}
	return User{} // Return zero value if not found
}

func Update(id int, updatedUser User) User {
	for i, user := range Users {
		if user.Id == id {
			Users[i] = updatedUser
			return updatedUser
		}
	}
	return User{} // Return zero value if not found
}

func Delete(id int) {
	var temp [] User;
	for _, user := range Users {
		if user.Id != id {

			temp=append(temp, user)
			
		}
	}
	Users = temp;
}