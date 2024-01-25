package service

type User struct {
	id	int32
	name 	string
	email 	string
	location Location
}

type Location struct {
	city	string
	state	string
		
}

func (u *User) getUserByID() {}
func (u *User) getUserByName() {}
func (u *User) getUserByEmail() {}
