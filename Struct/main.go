package main

import "fmt"

func main() {
	type person struct {
		name    string
		age     int
		contact string
	}
	p1 := person{
		"anu", 28, "8756432123",
	}
	p2 := person{
		name:    "jyoti",
		age:     27,
		contact: "8756432123",
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.name, p1.age, p1.contact)
	// **Difference Between %v and %#v
	// Verb	Behavior	Example Output
	// %v	Default formatting for the value.	{Alice 30} (for a struct)
	// %#v	Go syntax representation of value.	main.Person{Name:"Alice", Age:30}

	fmt.Printf("%T,%#v", p1, p1)

	// 	Embedded structs
	// We can take one struct and embed it in another struct. When you do this the inner type gets
	// promoted to the outer type.
// ******in go we can create field and method of employee type , 
// ******and wecan use these methods and fields from outer struct 
	type employee struct { 
		empid       int
		designation string
	}
	type engineer struct {
		team       string
		employee   //known as embedded struct, (method and variable are accessible to enginner type(like inheritance))
		desgnation string
	}

	e := engineer{
		"marketplace",
		employee{
			95,
			"system engineer",
		},
		"senior engineer",
	}

	fmt.Println(e)
	fmt.Println(e.desgnation, e.employee)
	fmt.Println(e.employee.designation)
	type staff struct {
		team string
		emp  employee //we can also decalre variable of type employee
	}
	st := staff{
		team: "product",
		emp: employee{
			empid:       1,
			designation: "manager",
		},
	}

	fmt.Println("******", st.emp) //Promoted fields empid & designation

	// Anonymous structs
	// We can create anonymous structs also. An anonymous struct is a struct which is not
	// associated with a specific identifier

	a := struct { //no need to decalre identifier only create value directly
		name string
		age  int
	}{
		"payal",
		25,
	}
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)

	// Pointers to a struct
	type s struct {
		name string
	}

	s1 := &s{
		name: "anu",
	}
	fmt.Println((*s1).name)
	fmt.Println(s1.name) // The Go language gives us the option to use s1.name instead of the explicit dereference.

	//Anonymous fields
	type p struct {
		string
		int
	}
	p3 := p{
		string: "anu", //can happen only in go
		int:    5,
	}

	fmt.Println(p3.int)

}
