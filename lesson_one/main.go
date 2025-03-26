package main

import "fmt"

func main () {

	s := "Müller"
	fmt.Printf("\n----EXAMPLE 1----\n")
	fmt.Printf("%T %[1]v\n", s)
	fmt.Printf("%8T %[1]v\n", []rune(s))
	fmt.Printf("%8T %[1]v\n", []byte(s))

	fmt.Printf("\n----EXAMPLE 2----\n")
	fmt.Printf("ranging\n")
	for index, r := range s {
		fmt.Printf("character:%c index:%d\n", r, index)
	}

	fmt.Printf("\n----EXAMPLE 3----\n")
	fmt.Printf("indexing\n")
	for i := 0; i < len(s); i++ {
		fmt.Printf("character:%c index:%d \n", s[i], i)
	}

 
}
