package main

import (
	"fmt"
	. "my-github/local-memory-cache-golang/service"
)

type Val struct {
	value string
}

func main() {

	lr := NewLRUConnection()
	cc := NewCCConn()

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

	cc.Set("A", v)

	fmt.Println(lr.Get("A"), lr.Get("B"), lr.Get("C"), lr.Length(), lr.Keys())
	fmt.Println(cc.Get("A"))

}
