package main

import "fmt"

func main () {

	

	// stringExample()
	// stringSlicing()

	// array()
	// slice()
	mapExm()
}

func stringExample() {
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

func stringSlicing() {
	fmt.Printf("\n ----EXAMPLE 4----\n")
	s := "the quick brown fox"
	fmt.Printf("old_s: %v pointer:%v\n", s, &s)

	fmt.Printf("\n ----EXAMPLE 5----\n")
	a := len(s)
	fmt.Printf("a: %v\n", a)

	fmt.Printf("\n ----EXAMPLE 6----\n")
	b := s[:3]
	fmt.Printf("b: %v pointer:%v\n", b, &b)

	fmt.Printf("\n ----EXAMPLE 7----\n")
	с := s[4:9]
	fmt.Printf("c: %v\n", с)

	fmt.Printf("\n ----EXAMPLE 8----\n")
	d := s[:4] + "slow" + s[9:]
	fmt.Printf("d: %v\n", d)

	fmt.Printf("\n ----EXAMPLE 9----\n")
	s += "es"
	fmt.Printf("new_s: %v pointer: %v\n", s, &s)

}


func array() {
	fmt.Print("----EXAMPLE 1----\n")
	var a [3]int
	fmt.Printf("a: %v\n", a)

	fmt.Print("----EXAMPLE 2----\n")
	b := [3]int{0, 1, 0}
	fmt.Printf("b: %v\n", b)	

	fmt.Print("----EXAMPLE 3----\n")
	c := [...]int{0, 0, 0}
	fmt.Printf("c: %v\n", c)	

	fmt.Print("----EXAMPLE 4----\n")
	fmt.Printf("a == b: %v\n", a == b)
	fmt.Printf("a == b: %v\n", a == c)

	fmt.Print("----EXAMPLE 5----\n")
	var d [3]int
	d = b // copy elements!!!
	fmt.Printf("d: %v\n", d)

	fmt.Print("----EXAMPLE 6----\n")
	b[0] = 1
	fmt.Printf("b[0]:%v d[0]:%v\n", b[0], d[0])


	fmt.Print("----EXAMPLE 7----\n")
	m := [...]int{1, 2, 3, 4}
	fmt.Printf("m: %v\n", m)

	fmt.Print("----EXAMPLE 8----\n")
	fmt.Printf("a:%v\n", a)
	updateArray(a)
	fmt.Printf("a_updated:%v\n", a)
	
	

}

func updateArray(a [3]int) {
	a[0] = 100
}


func slice () {
	fmt.Printf("----Example 1 ----\n")
	var a []int
	fmt.Printf("a: %v\n", a)

	fmt.Printf("----Example 2 ----\n")
	var b = []int{1, 2}
	fmt.Printf("b: %v\n", b)

	fmt.Printf("----Example 3 ----\n")
	var c = make([]int, 0)
	fmt.Printf("c: %v\n", c)

	fmt.Printf("----Example 4 ----\n")
	c = append(c, 1)
	fmt.Printf("c: %v\n", c)

	fmt.Printf("----Example 6 ----\n")
	a = append(a, 1)
	fmt.Printf("a: %v\n", a)

	fmt.Printf("----Example 7 ----\n")
	b = append(b, 3)
	fmt.Printf("b: %v\n", b)

	a = b

	fmt.Printf("----Example 8 ----\n")
	d := make([]int, 5)
	fmt.Printf("d: %v capacity: %v\n", d, cap(d))


	fmt.Printf("----Example 9 ----\n")
	e := a
	fmt.Printf("e[0] == b[0]: %v\n", e[0] == b[0])

	fmt.Printf("----Example 10 ----\n")
	fmt.Printf("a:%v\n", a)
	updateSlice(a)
	fmt.Printf("a_updated:%v\n", a)


	fmt.Printf("----Example 11 ----\n")
	var oneMoreNil []int
	fmt.Printf("len one_more_nil: %v\n", len(oneMoreNil))
	fmt.Printf("cap one_more_nil: %v\n", cap(oneMoreNil))

	for index, value := range a {
		fmt.Printf("index:%v value:%v", index, value)
	}
	
}

func updateSlice(a []int) {
	a[0] = 100
}

func mapExm () {

	fmt.Printf("\n----Example 1----\n")
	var m map[string]int
	fmt.Printf("m:%v\n", m)

	fmt.Printf("\n----Example 2----\n")
	p := make(map[string]int)
	fmt.Printf("p:%v\n", p)
	
	fmt.Printf("\n----Example 3----\n")
	fmt.Printf("m == nil:%v\n", m == nil)
	fmt.Printf("p == nil:%v\n", p == nil)

	fmt.Printf("\n----Example 4----\n")
	a := p["the"]
	fmt.Printf("a:%v\n", a)

	fmt.Printf("\n----Example 5----\n")
	b := m["the"]
	fmt.Printf("b:%v\n", b)


	fmt.Printf("\n----Example 6----\n")
	// m["and"] = 1
	p["smth"] = 1

	m = p

	m["and"]++
	c := p["and"]
	fmt.Print(m, p, c)
	fmt.Printf("c:%v\n", c)


	fmt.Printf("\n----Example 7----\n")
	or := p["or"]
	fmt.Printf("or:%v\n", or)

	fmt.Printf("\n----Example 8----\n")
	not, ok := p["not"]
	fmt.Printf("not:%v ok:%v\n", not, ok)


	fmt.Printf("\n----Example 9----\n")
	if not, ok := p["not"]; ok {
		fmt.Printf("existing not: %v\n", not)
	}

	fmt.Printf("\n----Example 10----\n")
	for key, value := range p {
		fmt.Printf("key:%v value:%v\n", key, value)
	}



	fmt.Printf("\n----Example 11----\n")
	sortedMap := map[int]string {
		100: "a",
		1: "b",
		-3: "c",
	}

	fmt.Printf("sorted:%v\n", sortedMap)

}