package main

import (
	fmt "fmt"
	"sort"
	rstring "workshop/pack"
)

type A struct {
	a int
}

type B struct {
	b int
	a int
	A
}

func printA(a A) {
	fmt.Printf("%v\n", a)
}

func main() {
	b := B{1, 3, A{2}}
	printA(b.A)

	rs := rstring.NewRString("hello world")
	rs1 := rstring.NewRString("a")
	rs2 := rstring.NewRString("ab")
	rs3 := rstring.NewRString("abc")

	ary := rstring.SortableRString{rs, rs1, rs3, rs2}
	fmt.Printf("%v\n", ary)
	sort.Sort(ary)
	fmt.Printf("%v\n", ary)
}
