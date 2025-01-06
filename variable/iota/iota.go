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
	fmt.Printf("Decimal value = %d, Binary value = %b\n",a,a)
	fmt.Printf("Decimal value = %d, Binary value = %b\n",b,b)
	fmt.Printf("Decimal value = %d, Binary value = %b\n",c,c)
	fmt.Printf("Decimal value = %d, Binary value = %b\n",d,d)
	fmt.Printf("Decimal value = %d, Binary value = %b\n",e,e)
	// create a program that uses iota to calculate the size of each measurement of bytes
	// ○ KB 1024 bytes
	// ○ MB 1024 KB
	// ○ GB 1024 MB
	// ○ TB 1024 GB
	// ○ PB 1024 TB
	// ○ EB 1024 EB
	const(
		_ = iota
		KB = 1<<(10*iota)
		MB //= 1<<(10*iota)- can omit this part, coz iota now takes care of rest.
		GB //= 1<<(10*iota)
		TB //= 1<<(10*iota)
		PB //= 1<<(10*iota)
		EB//= 1<<(10*iota)

	)
	fmt.Printf("Decimal value = %d, Binary value = %b\n",KB,KB)
	fmt.Printf("Decimal value = %d, Binary value = %b\n",MB,MB)
	fmt.Printf("Decimal value = %d, Binary value = %b\n",GB,GB)
	fmt.Printf("Decimal value = %d, Binary value = %b\n",TB,TB)
	fmt.Printf("Decimal value = %d, Binary value = %b\n",PB,PB)
	fmt.Printf("Decimal value = %d, Binary value = %b\n",EB,EB)
}