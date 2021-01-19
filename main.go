package main

import "fmt"

type Foo struct {
}

func (f *Foo) String() string {
  return "foo"
}

func main() {
  fmt.Printf("Hello\n")
  f := &Foo{}
  fmt.Printf("Hello %+v\n", f)
}
