package main

import (
	"fmt"
	. "my-github/local-memory-cache-using-lru/service"
)

type Val struct {
	value string
}

func main() {

	lr := NewLRUConnection()

	// var s, t, u interface{}
	v := &Val{
		value: "50",
	}
	// x := &Val{
	// 	value: "100",
	// }
	// var xv []interface{}
	// xv = append(xv, *v, *x)
	lr.Set("A", v)
	// lr.Set("A", x)
	// lr.Get("A", &s)
	lr.Set("B", v)
	// lr.Get("B", &t)
	lr.Set("C", v)
	// lr.Get("C", &u)

	fmt.Println(lr.Get("A"), lr.Get("B"), lr.Get("C"), lr.Length(), lr.Keys())

}
