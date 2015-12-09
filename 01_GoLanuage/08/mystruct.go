package mystruct

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
	age  int
}

func (u *User) SayHi() {
	fmt.Printf("Hi, I'm %s. Nice to meet you!\n", u.Name)
}

func (u User) old() int {
	return u.age
}
