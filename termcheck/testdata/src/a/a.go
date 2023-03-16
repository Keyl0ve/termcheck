package a

import (
	"user"
	uu "user"
)

// パッケージと関数の term check
func checkCallFunction() {
	// Bad
	user.ReadUser()          // want "word is used multiple in same line"
	user.ReadUserFromJapan() // want "word is used multiple in same line"
	// Good
	uu.ReadUser()          // OK
	uu.ReadUserFromJapan() // OK
}

type User struct {
	userName string
}

// 構造体の term check
// Good
func (u User) a() {
	u.userName = "aaa" // OK
}

// Bad
func (user User) b() {
	user.userName = "aaa" // want "word is used multiple in same line"
}
