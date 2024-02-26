package main

import "fmt"

func main() {
	fmt.Println("Sai ram🙏")
	
	var a int
	fmt.Println(a)
// ***********statically typed************
	a=43
	//a="jyoti" //you cannt change type means it is statically type not dynamically type.

	//a:=5 //you cannt re-initialized variable identifier
	fmt.Println(a)

	//a:="jyoti" //you cannt re-initialized variable identifier for different type.
	// ***********declaration************
	g:="jyoti"
	fmt.Println(g)

	c:=44
	fmt.Println(c)
	c=45
	fmt.Println(c)
	//c=56.7 //cannt use float type

	//b, c, d, _, f := 0, 1, 2, 3, "happiness" //if dont want to print then use blank indentifier
	//fmt.Println(b, c, d, f)
// ***********housekeeping************
	// this would not work
	/* coz e is decalre and intialize but not used
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d)
	*/
// ***********type************
	// b,c,d:=1,2,3
	// fmt.Printf("value in binary %b %b %b \n",b,c,d)
	// fmt.Printf("value in hexadecimal %x %x %x \n",b,c,d)

	// ***********conversion************

	y := 42
	z := 42.0
	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", m, m)

	/*
		// this does not work!
		// in go you can't take a VALUE that is float32 and store it 
		// in a variable that is declared to hold a VALUE of float64
		z = m 
		fmt.Printf("%v of type %T \n", z, z)
	*/

	
		// this does work
		z = float64(m)
		fmt.Printf("%v of type %T \n", z, z)
	

	// ***********************

}