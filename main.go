package main

import (
	"reflect"
	"sync/atomic"
)

func main() {
	x := 2                   // value   type    variable?
	a := reflect.ValueOf(2)  // 2       int     no
	b := reflect.ValueOf(x)  // 2       int     no
	c := reflect.ValueOf(&x) // &x      *int    no
	d := c.Elem()            // 2       int     yes (x)

	var z atomic.Bool
	z.CompareAndSwap()
}
