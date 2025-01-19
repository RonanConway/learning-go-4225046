package main

import (
	"fmt"
)

func main() {
	
	anInt := 42
	var p *int = &anInt

	if p == nil{
		fmt.Println(" p is nil")
	} else {
	  fmt.Println("Value of P : ", *p)
	}


	value1 :=42.13
	pointer := &value1
	*pointer = *pointer / 31
	fmt.Println("Value 1 : ", *pointer)
	fmt.Println("Original Value : ", value1)

}
