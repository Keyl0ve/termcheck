package a

import (
	"user"
	uu "user"
)

// パッケージと関数の term check
func checkCallFunction() {
	// Bad
	user.ReadUser()
	// Good
	uu.ReadUser() // OK
}

type User struct {
	userName string
}

// 構造体の term check
// Good
func (u User) a() {
	u.userName = "aaa"
}

// Bad
func (user User) b() {
	user.userName = "aaa"
}
