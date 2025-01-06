package main

import (
	"math/rand"
	"fmt"
	//"time"
)
func init(){ //this is known as niladic function means which will take no argument and no parameter
	fmt.Println("this is called before main function")
}
func Random()(int){
	return rand.Intn(250)
}
func main() {
// 	a := 40
// 	//b := 45

// // ******logical operation (&&) , comparison (>) and if block *************
// 	// if a > 40 && a < 45 { 
// 	// 	fmt.Printf("value of a=%d",a)
// 	// } else{
// 	// 	fmt.Printf("value of b = %d",b)
// 	// }
// // **************switch statement ****************

// switch{
// case a>40 :
// 	fmt.Println("value a is greater")
// case a<45:
// 	fmt.Println("value a is smaller")
// 	fallthrough
// case a==40:
// 	fmt.Println("value a is exact same")
// default:
// 	fmt.Println("print default case" )
// }

// switch a{
// case 40:
// 	fmt.Println("40")
// 	fallthrough
// case 41:
// 	fmt.Println("41")
// case 42:
// 	fmt.Println("42")
// default:
// 	fmt.Println("default value")
// }

// // *************select statement***********
//  c1 := make(chan string)
//  c2:= make(chan string)
// go func(){
// 	time.Sleep(1*time.Second)
// 	c1 <- "one"
// }()
// go func(){
// 	time.Sleep(2*time.Second)
// 	c1 <- "two"
// }()
// // no loop so select will run one time and also here we are forcing to run one first as only 1 sec pause
// // Each channel will receive a value after some amount of time, to simulate here blocking RPC operations for given time
// // to executing in concurrent goroutines.
// // ***it will print one ***
// // select{
// // case c1:= <- c1:
// // 	fmt.Println(c1)
// // case c2:= <-c2:
// // 	fmt.Println(c2)
// // }
// // ***here it will print all go rountine output coz of for loop****
// for{
// select{
// case c1:= <- c1:
// 	fmt.Println(c1)
// case c2:= <-c2:
// 	fmt.Println(c2)
// }
// }

// ****Exercise 19+20 *****
x:=Random()
switch {
case x<=100:
	fmt.Printf("random value of x b/w 0 to 100 is %v\n",x)
case x>=101 && x<=200 : 
	fmt.Printf("random value of x b/w 101 to 200 is %v\n",x)
case x>=250:
		fmt.Printf("random value of x b/w 201 to 250 is %v\n",x)
default : 
fmt.Println("invalid\n")
}
// ****Exercise 24 *****
// var i int
// for i=0;i<=100;i++ {
	// fmt.Printf("%d",i)
// }
// ****Exercise 25 *****
y:=rand.Intn(5)
switch y {
case 0:
	fmt.Println("0")
case 1:
	fmt.Println("1")
case 2:
	fmt.Println("2")
case 3:
	fmt.Println("3")
case 4:
	fmt.Println("4")
default:
	fmt.Println("no new value/this should not print")
}
c:=1
// for c<5 { //can be written as for ; c<5 ; {}
// fmt.Println("ok to proceed")
// c++
// }
for{
	if c==5{
		break
	}
	c++// without giving increment this loop will run infinite loop
	fmt.Println("ok to proceed")
}
// ****Exercise 35 +36+37 *****
// xi := []int{42, 43, 44, 45, 46, 47} // slice
m := map[string]int{ //map
	"James": 42,
	"Moneypenny": 32,
	}
for key,value:=range m{
	if(key=="Q"){
		fmt.Println("yes")
	}else{
		fmt.Println("No")
	}
	fmt.Printf("value at index %s is %d\n",key,value)
}
//comma ok
if value,ok:=m["Q"]; !ok{
fmt.Println("comma ok",value)
}
//statement statement
// stmt:=rand.Intn(5)
i:=0
c=0
for i<100{
if y=rand.Intn(5);y==3{
	fmt.Println("its 3")
	c++
} 
i++
}
fmt.Printf("total count of 3 is %d\n",c)
age:=m["James"]
fmt.Println(age)
}