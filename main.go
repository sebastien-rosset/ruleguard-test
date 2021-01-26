package main

import (
  "fmt"
  "github.com/sebastien-rosset/rg1"
)

type Foo struct {
}

type Bar struct {
}

func (f *Foo) String() string {
  return "foo"
}

func (f *Bar) Run() {
  fmt.Printf("Running\n")
}

func main() {
  fmt.Printf("Hello\n")
  f := &Foo{}
  fmt.Printf("Stringer %+v\n", f)

  var w rg1.Worker  = &Bar{}
  w.Run()
}
