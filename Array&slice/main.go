package main

import "fmt"

func change(x []int) []int{
	x[0]=10
	return x
}

func main() {
	var a [2]int //You have to populate array size else without declare size it will give runtime error.
	a[0] = 1
	a[1] = 2
	fmt.Printf("%d,%d,%d\n", a[0], a[1], len(a))
	b := [2]string{"jyoti", "anu"}
	fmt.Println(b[0])
	c := []float32{2.5, 3.5} //here not decalre array ize directly but indirectly declare
	fmt.Println(c[0], c[1], len(c))
	var d [3]int
	fmt.Println(len(d))

	fmt.Printf("%T\n%T\n%T\n", a, b, d)
	// a=b //error because you can not assign different-2 type of array
	// a=d // error because here both are different type (length is different)

	// ***Exercise 40 ***
	// e:=[]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	// fmt.Println(len(e))
	// for _,v:=range e{
	// 	fmt.Println(v)
	// }
	// e=append(e, "jyoti")
	// fmt.Println(len(e))
	//fmt.Println(e)
	// *****slicing a slice ****
	f := []int{3, 4, 5, 6, 7, 8, 9}
	fslice1 := f[1:4] //[inclusive:exclusive]
	fmt.Println(fslice1)
	fslice2 := f[1:] //[inclusive:]
	fmt.Println(fslice2)
	fslice3 := f[:4] //[:exclusive]
	fmt.Println(fslice3)
	fslice4 := f[:] //[:]
	fmt.Println(fslice4)

	// *****deleting from a slice ****

	fslice5 := append(f[:2], f[3:]...) //delete can happen only with append and slice , ... it is necessary for appending slice at the end.
	fmt.Println(fslice5)

	// *****create slice using make function ****
	fslice6 := make([]int, 1, 10)
	fmt.Println(fslice6)      //length 0 ,means empty slice
	fmt.Println(fslice6[0])   //accesing a empty slice has by defalut vlue 0 value
	fmt.Println(len(fslice6)) //length is size of slice length
	fmt.Println(cap(fslice6)) // capacity is size of underlying array

	//appending in an slice double the capacity if you are inserting out of size of underlying array
	fslice6 = append(fslice6, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(fslice6)
	fmt.Println(len(fslice6))
	fmt.Println(cap(fslice6)) // size not double coz leng of slice is same as capacity

	fslice6 = append(fslice6, 10)
	fmt.Println(fslice6)
	fmt.Println(len(fslice6))
	fmt.Println(cap(fslice6)) // here double coz cap is only 10 and appending 11th value
	// here it will craete different underlying array of new size . and capacity double at run time.

	// *****understanding od underlying array ****

	f1 := []int{1, 2, 3}
	f2 := f1
	fmt.Println(f1)
	fmt.Println(f2)
	f3 := make([]int, 3) //ev
	fmt.Println(f3)
	f1 = append(f1, 4)
	f3 = f1 //it means copy the value and underlying array too so changing in f1 will changein f3 too, eventhough it is craeted new
	f1[0] = 5
	fmt.Println(f3) //[5 2 3 4]

	//now to have different underlying array and wanna copy all element
	//first create new slice
	f4 := make([]int, 4)
	copy(f4, f1) //copy will pont to different underlying array but assgning is pointing to same underlying array
	fmt.Println(f1)
	fmt.Println(f4)
	f1[1] = 6
	fmt.Println(f1)
	fmt.Println(f4) //here change not happened in f4 coz of copy function


	// *** pass by value ***
	// everything in go in by default pass by value and pass by value means in go is assigning (means poining to same underlying array so changes in function will change in main function too)

	fmt.Println(f1)
	fmt.Println(change(f1)) //we can see changes in function so change here too, if dont want to change then create new alice and copy function will use means new underlying array.

}

