package a

import (
	"user"
	uu "user"
	uuu "user"
)

// パッケージと関数の term check
func checkCallFunction() {
	// Bad
	user.ReadUser()          // want "word is used multiple in same line"
	user.ReadUserFromJapan() // want "word is used multiple in same line"
	// Good
	uu.ReadUser()          // OK
	uu.ReadUserFromJapan() // OK

	uuu.ReadUser()          // OK
	uuu.ReadUserFromJapan() //OK
}

type User struct {
	userName string
	age      int
}

// 構造体の term check
// Good
func (u User) a() {
	u.userName = "aaa" // OK
	u.age = 10         // OK
}

func (uu User) b() {
	uu.userName = "aaa" // OK
	uu.age = 10         // OK
}

// Bad
func (user User) c() {
	user.userName = "aaa" // want "word is used multiple in same line"
}
