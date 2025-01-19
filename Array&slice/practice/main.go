package main

import "fmt"

func main() {
	fmt.Println("only one main file sould be present in one folder")

	/*
	Using a COMPOSITE LITERAL:
● create an ARRAY which holds 5 VALUES of TYPE int
● assign VALUES to each index position.
● Range over the array and print the values out.
○ print out the VALUE and the TYPE*/
	arr:=[5]int {}
	arr[0]=0
	arr[1]=1
	arr[2]=2
	arr[3]=3
	arr[4]=4

	for i,v:=range arr{
		fmt.Printf("index=%d,value=%d,and Type=%T\n",i,v,v)
	}

	/*● Using a COMPOSITE LITERAL:
● create a SLICE of TYPE int
● assign these 10 VALUES
42, 43, 44, 45, 46, 47, 48, 49, 50, 51
● Range over the slice and print the values out.
○ print out the VALUE and the TYPE*/ 
// -- similar to above 
slice:=[]int {42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// slice:=make([]int,10,10)
// for i:=0;i<10;i++{
// 	slice[i]= arr1[i]
// }
for _,v:=range slice{
	fmt.Printf("Value = %d, Type=%T \n",v,v)
}
/*
Using the code from the previous example, use SLICING to create the following new slices
which are then printed:
● [42 43 44 45 46]
● [47 48 49 50 51]
● [44 45 46 47 48]
● [43 44 45 46 47]*/

// slice1:=[4][5]int {{42, 43, 44, 45, 46},{47, 48, 49, 50, 51},{44, 45, 46, 47, 48},{43, 44, 45, 46, 47}}

// for _,v:=range slice1{
// 	fmt.Println(v)
// }
// [inclusive:exclusive]
fmt.Println(slice[0:5])
fmt.Println(slice[5:10])
// slice=append(slice, )
fmt.Println(slice[2:7])
fmt.Println(slice[1:6])
// fmt.Println(cap(slice1))

/*
Follow these steps:
● start with this slice
○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
● append to that slice this value
○ 52
● print out the slice
● in ONE STATEMENT append to that slice these values
○ 53
○ 54
○ 55
● print out the slice
● append to the slice this slice
○ y := []int{56, 57, 58, 59, 60}
● print out the slice
*/
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
x=append(x, 52)
fmt.Println(x)
x=append(x, 53,54,55)
fmt.Println(x)
y := []int{56, 57, 58, 59, 60}
x=append(x, y...)
fmt.Println(x)

/*
● start with this slice
○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
● use APPEND & SLICING to get these values here which you should ASSIGN to the
variable “x” and then print:
○ [42, 43, 44, 48, 49, 50, 51]
*/
x1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
x1 = slice[:3] 
x1=append(x1, slice[6:10]...)
fmt.Println(x1)

/*
For this exercise, do the following:
● Create a slice to store the names of all of the states in the United States of America.
○ Use make and append to do this.
○ Goal: do not have the array that underlies the slice created more than once.
● Print out
○ the len
○ the cap
○ the values, along with their index position, without using the range clause.
● Here is a list of the 50 states:
` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
*/

s:=make([]string,50)
fmt.Println(len(s))
fmt.Println(cap(s))
s=append(s, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

fmt.Println(len(s))
fmt.Println(cap(s))//112- coz of doublig the size and extra buffer depend upon version and go runtime
for i:=0;i<len(s);i++{
fmt.Printf("index=%d,value=%s\n",i,s[i])
}


fmt.Println("_________________________________________")
    s2 := make([]int, 0, 4) // Start with capacity 4
    for i := 0; i < 100; i++ {
        s2 = append(s2, i)
        fmt.Printf("Length: %d, Capacity: %d\n", len(s2), cap(s2))
    }


	/*
	Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
slice:
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."
Range over the records, then range over the data in each record
	*/

	s3:=[][]string {{"James", "Bond", "Shaken, not stirred"},{"Miss", "Moneypenny", "I'm 008."}}
	for _,v:=range s3{
		fmt.Println(v)
	}
}