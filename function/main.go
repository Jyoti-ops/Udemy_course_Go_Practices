package main

import "fmt"

func main() {
//defer stmt
Abc()
Bcd("Jyots")
s:=Cde("jyoti")
fmt.Println(s)
s,s1:=Def("jyoti",25)
defer fmt.Println(s,s1) // this stmt will execute last when function returns
a,b:=Efg(15,20)
fmt.Println(a,b)
p:=person{
	"jyoti",
}
p.call()
}

// no params, no returns
// ○ 1 param, no returns
// ○ 1 param, 1 return
// ○ 2 params, 2 returns

func Abc() {
	fmt.Println("No param, no return")
}

func Bcd(s string){
	fmt.Println("1 param and no return")
}

func Cde(s string)(string){
return fmt.Sprint("1 param and 1 return , ok ",s)
}

func Def(name string,age int)(string, int){
	return name,age
}

func Efg(a,b int)(b1,a1 int){
	return b,a
}

//method

// when a function attched to a type then this function named as method
type person struct{
	name string
}

func (p person) call(){
	fmt.Println(p.name)
}




