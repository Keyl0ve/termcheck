# termcheck
termcheck is the linter to check using simple term

## 概要
↓の時に同じ単語を使っていたら警告を出す linter です。
- パッケージ名 . 関数名
- 構造体 . フィールド名

## Example

#### パッケージ名 . 関数名

```go
func a() {
	user.Read()              // OK
	user.ReadUser()          // want "user is used multiple in same line"
	user.ReadUserFromJapan() // want "user is used multiple in same line"

	uu.Read()              // OK
	uu.ReadUser()          // OK
	uu.ReadUserFromJapan() // OK
}
```
---

#### レシーバー . フィールド名


```go
func (u User) a() {
	u.userName = "John"          // OK
	u.age = 10                   // OK
	u.country.address = "Japan"  // OK
}

func (user User) a() {
	user.userName = "John"         // want "user is used multiple in same line"
	user.country.myNumber = 1234   // OK
	user.country.address = "Japan" // OK
}
```

## install
```sh
go install github.com/Keyl0ve/termcheck
```

## usage
```sh
go vet -vettool=$(which termcheck) ./...
```

### ref
- https://pkg.go.dev/golang.org/x/tools/go/analysis
- https://pkg.go.dev/golang.org/x/tools/go/ast/inspector
- https://pkg.go.dev/github.com/iancoleman/strcase@v0.2.0
- https://pkg.go.dev/strings
