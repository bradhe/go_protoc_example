package main

import (
  "errors"
  "os"
  "counters"
  "fmt"
)

func main() {
  service, err := counters.NewCountersService("localhost:3000")

  if err != nil {
    panic(err)
  }

  op := os.Args[1]
  name := os.Args[2]

  if op == "inc" {
    fmt.Printf("%s: %d\n", name, service.Increment(name))
  } else if op == "dec" {
    fmt.Printf("%s: %d\n", name, service.Decrement(name))
  } else if op == "get" {
    fmt.Printf("%s: %d\n", name, service.Get(name))
  } else {
    panic(errors.New(fmt.Sprintf("Unsupported operation %s.", op)))
  }
}
