package a

import (
	"fmt"
	"user"
	uu "user"
	uuu "user"
	uuuu "user"
)

// パッケージと関数の term check
func checkCallFunction() {
	user.Read()                 // OK
	aaa := user.ReadUser("aaa") // want "user is used multiple in same line"
	fmt.Println(aaa)
	user.ReadUserFromJapan() // want "user is used multiple in same line"

	uu.Read()              // OK
	uu.ReadUserFromJapan() // OK

	uuu.Read()              // OK
	uuu.ReadUserFromJapan() // OK

	uuuu.Read()              // OK
	uuuu.ReadUserFromJapan() // OK
}

type User struct {
	userName string
	age      int
	country  Country
}

type Country struct {
	address    string
	userNumber string
	userFriend Friend
}

type Friend struct {
	userNum int
}

// 構造体の term check
func (u User) a() {
	u.userName = "John"          // OK
	u.age = 10                   // OK
	u.country.address = "Japan"  // OK
	u.country.userNumber = "aaa" // OK
}

func (uu User) b() {
	uu.userName = "aaa"                // OK
	uu.age = 10                        // OK
	uu.country.userFriend.userNum = 10 // want "user is used multiple in same line"
}

func (uuu User) c() {
	uuu.userName = "aaa" // OK
	uuu.age = 10         // OK
}

func (uuuu User) d() {
	uuuu.userName = "aaa" // OK
	uuuu.age = 10         // OK
}

func (user User) e() {
	user.userName = "John"          // want "user is used multiple in same line"
	user.country.userNumber = "aaa" // want "user is used multiple in same line"
	user.country.address = "Japan"  // OK
}
