package main

import "fmt"

type A struct {
	balance   int
	firstname string
	lastname  string
}

func (a1 *A) widthdrawPointer(amount int) {
	a1.balance -= amount
}

func (a1 A) widthdrawValue(amount int) {
	a1.balance -= amount
}

type myInt int

func (m myInt) Add(a int) myInt {
	rst := int(m) + a
	return myInt(rst)
}

type account struct {
	balance int
}

func widthdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// 리시버
func (a *account) widthdrawMethod(amount int) {
	a.balance -= amount
}

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	//Sprintf => string 값을 만들어 반환함
	return fmt.Sprintf("%s %d", u.Name, u.Age)
}

type VIPuser struct {
	User
	VIPleve int
}

func main() {
	//a=> account pointer
	a := &account{100}

	widthdrawFunc(a, 30)

	a.widthdrawMethod(30)

	var mainA *A = &A{100, "A", "B"}
	mainA.widthdrawPointer(30)
	fmt.Println(mainA.balance)

	mainA.widthdrawValue(20)
	fmt.Println(mainA.balance)

	vip := VIPuser{User{"A", 34}, 5}
	fmt.Println(vip.String())
}
