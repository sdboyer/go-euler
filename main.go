package main

// My stab at some Project Euler problems.

import (
	"flag"
	"fmt"
	"reflect"
)

var fml = fmt.Println

var q = flag.String("q", "", "The question to run")

type Euler uintptr

func main() {
	flag.Parse()
	if *q == "" {
		panic("must provide a question number")
	}

	var e Euler
	reflect.ValueOf(e).MethodByName("Problem" + *q).Call(make([]reflect.Value,0))
}
