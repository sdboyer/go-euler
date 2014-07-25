package main

// My stab at some Project Euler problems.


import (
	"flag"
	"fmt"
)

var fml = fmt.Println

var q = flag.Int("q", 0, "The question to run")

var answers = make(map[int]func(), 0)

func main() {
  flag.Parse()
  if *q == 0 {
    panic("must provide a question number")
  }

  answers[*q]()
}

