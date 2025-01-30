package main

import "fmt"

type person struct {
	firtsname       string
	lastname        string
	icecreamflavour []string
}

func main() {
	/*
			Create your own type “person” which will have an underlying type of “struct” so that it can store
		the following data:
		● first name
		● last name
		● favorite ice cream flavors
		Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
		which stores the favorite flavors.
	*/

	Student := person{
		firtsname:       "anu",
		lastname:        "anand",
		icecreamflavour: []string{"chocolate", "strawberry"},
	}
	baby := person{
		firtsname:       "JYOIT",
		lastname:        "JYOTS",
		icecreamflavour: []string{"chocolate", "strawberry", "VANILLA"},
	}
	// s := []string{Student.icecreamflavour, baby.icecreamflavour}
	fmt.Println(Student)
	fmt.Println(baby)
	for _, v := range baby.icecreamflavour {
		fmt.Println(v)
	}

	/*
		Take the code from the previous exercise, then store the VALUES of type person in a map with
	the KEY of last name. Access each value in the map. Print out the values, ranging over the
	slice.
	*/
	fmt.Println("****************")
	m := map[string]person{
		baby.lastname: baby,
	}

	for i, v := range m {
		fmt.Println(i, v.icecreamflavour)
	}

	/*
	   Create a type engine struct, and include this field
	   ○ electric bool
	   ● Create a type vehicle struct, and include these fields
	   ■ engine
	   ■ make
	   ■ model
	   ■ doors
	   ■ color
	   ● Create two VALUES of TYPE vehicle
	   ○ use a composite literal
	   ● Print out each of these values.
	   ● Print out a single field from each of these values
	*/

	fmt.Println("&&&&&&&&&&")

	type engine struct {
		electric bool
	}
	type vehicle struct {
		e     engine
		make  string
		model int
		door  string
		color string
	}
	V1 := vehicle{
		e: engine{
			electric: true,
		},
		make:  "yes",
		model: 5,
	}
	v2 := vehicle{
		e: engine{
			false,
		},
		color: "black",
	}

	fmt.Println(V1)
	fmt.Println(v2)
	fmt.Println(v2.e)
	fmt.Println(v2.e.electric)

	/*
	   Create and use an anonymous struct with these fields:
	   ● first string
	   ● friends map[string]int
	   ● favDrinks []string
	   Print things.
	*/
	fmt.Println("#############")
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "anu",
		friends: map[string]int{			//map initialize as it is initializing 
			"khusi": 25,
			"jyoti": 24,
		},
		favDrinks: []string{"coke", "lemon juice"},
	}

	fmt.Println(s)

}
