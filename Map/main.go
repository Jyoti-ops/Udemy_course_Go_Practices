package main

import "fmt"

func main() {
	m := map[string]int{
		"abc": 5,
		"def": 6,
		"ghi": 7,
	}
	fmt.Println(m)
	//Map - for range over a map
	for i, v := range m {
		fmt.Println(i, v) //value will print at random
	}
	//Map - delete element
	fmt.Println(m)
	delete(m, "ghi")
	fmt.Println(m)
	delete(m, "jkl") //not exist in map but no runtime error
	fmt.Println(m)

	//Map - comma ok idiom
// If you look up a non-existent key, the zero value will be returned as the value associated
// with that non-existent key. You can check to see if you have looked up an existing key, or a
// non-existent key, using the comma ok idiom
v,ok:=m["abc"] //comma-ok
if ok{
	fmt.Println(v)
}else{
	fmt.Println("not found")
}
fmt.Println("**********")
if v,ok:=m["abc"];ok{ //stmt;stmt
	fmt.Println(v)
}else{
	fmt.Println("not found")
}
fmt.Println("**********")
if v,ok:=m["jkl"];ok{ //stmt;stmt , key not exist and we can access there will be no runtime error
	fmt.Println(v)
}else{
	fmt.Println("not found")
}

fmt.Println("*********")

// text:="Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. Go is expressive, concise, clean, and efficient."
// m3:=make(map[string]int)

// for i,v:=range m{
// 	m[i]++  //count word 
// }

}
