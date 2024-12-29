package iota

import "fmt"

func Iota() {
	const (
		_ = iota
		a
		b
		c
		d
		e
	)
	fmt.Println("Decimal value = %d, Binary value = %b",a,a)
	fmt.Println("Decimal value = %d, Binary value = %b",b,b)
	fmt.Println("Decimal value = %d, Binary value = %b",c,c)
	fmt.Println("Decimal value = %d, Binary value = %b",d,d)
	fmt.Println("Decimal value = %d, Binary value = %b",e,e)
}