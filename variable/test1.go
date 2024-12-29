package main

import (
"fmt"
"Udemy_course_Go_Practices/variable/iota"
)

func main() {
	fmt.Println("Sai ram🙏")
	iota.Iota()
	var a int
	fmt.Println(a)
// ***********statically typed************
	a=43
	//a="jyoti" //you cannt change type means it is statically type not dynamically type.

	//a:=5 //you cannt re-initialized variable identifier
	fmt.Println(a) 

	
	//a="jyoti" //you cannt initialized variable identifier for different type.
	// ***********declaration************
	g:="jyoti"
	fmt.Println(g)

	c:=44
	fmt.Println(c)
	c=45
	fmt.Println(c)
	//c=56.7 //cannt use float type
// ***********multiple declaration************
	b, c, d, _, f := 0, 1, 2, 3, "happiness" //if dont want to print then use blank indentifier
	fmt.Println(b, c, d, f)
// ***********housekeeping************
// housekeeping means : It refers management and maintenance of a household, which includes tasks like cleaning, dusting, laundry, and decluttering.
// 1. decalre but not used
// 2. short declaration only be declared inside method , would not be used as global value decalaration
// 3. we cannt initialized variable with other types ex :  var f float32 ; var f1 float64 ; f=f1 -  it cannt be happen beacuse both are different type.
// 4. It is an error to import a package or to declare a variable without using it - resolved using blank identifier
	// this would not work
	/* coz e is decalre and intialize but not used
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d)
	*/
// ***********type************
	o,n,p:=1,2,3
	fmt.Printf("value in binary %b %b %b \n",o,n,p)
	fmt.Printf("value in hexadecimal %x %x %x \n",o,n,p)
	fmt.Printf("value in hexadecimal %#x %#x %#x \n",o,n,p) // here more like to represent in hexadecimal

	// ***********conversion************
// An explicit conversion is an expression of the form T(x) where T is a type and x is an
// expression that can be converted to type T.
	y := 42
	z := 42.0
	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", m, m)

	/*
		this does not work!
		in go you can't take a VALUE that is float32 and store it 
		in a variable that is declared to hold a VALUE of float64
		z = m 
		fmt.Printf("%v of type %T \n", z, z)
	*/

	
		// this does work
		z = float64(m)
		fmt.Printf("%v of type %T \n", z, z)
	// ***********Constants************
		// Constants in Go are just that—constant. They are created at *compile time*, even when
		// defined as locals in functions, and can only be numbers, characters (runes), strings or
		// booleans. Because of the compile-time restriction, the expressions that define them must
		// be constant expressions, evaluatable by the compiler. For instance, 1<<3 is a constant
		// expression, while math.Sin(math.Pi/4) is not because the function call to math.Sin needs
		// to happen at run time. Constants may be typed or untyped.Go is a statically typed language that does not permit operations that mix numeric types. You can’t add a float64 to an int, or even an int32 to an int. In Go, constants, unlike variables,behave pretty much like regular numbers.
		const k="jyoti"
		const k1 = "abc"
		// const k2 int // must initialize value at the time of decaration unlike variable as it should be compute at compile time.
		// k2=5 // cannt d for constant type value
		// k3:=8// for const we should declare as const otherwise it would not recognize as constant.means const cannt be declare with shorthand. 
	// ***********intialize with zero************
	var s string //by default initialize with blank coz of string type
	var i int //by default initialize with blank coz of int type
	fmt.Printf("value of string and int = %s,%d",s,i)
	// general guideline

	// 	○ var for zero value - declare with var only need to print default value ie 0/ " "
	// 	○ short declaration operator -- in general,  mostly use this
	// 	○ var when specificity is required - declare when reuired

}