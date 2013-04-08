package main

import (
  "counters"
  "log"
  "fmt"
)

var counts = make(map[string] int32)

type countersServiceImpl struct {
  // TODO: Add...anything...
}

// Implementation of the Increment interface function.
func (c countersServiceImpl) Increment(in *counters.CountRequest, out *counters.CountResponse) error {
  counts[in.GetName()] += 1
  out.SetCount(counts[in.GetName()])

  log.Println(fmt.Sprintf("Incrementing %s to %d", in.GetName(), counts[in.GetName()]))

  return nil
}

// Implementation of the Decrement interface function.
func (c countersServiceImpl) Decrement(in *counters.CountRequest, out *counters.CountResponse) error {
  counts[in.GetName()] -= 1

  if counts[in.GetName()] < 0 {
    counts[in.GetName()] = 0
  }

  out.SetCount(counts[in.GetName()])

  log.Println(fmt.Sprintf("Decrementing %s to %d", in.GetName(), counts[in.GetName()]))

  return nil
}

// Implementation of the Get interface function.
func (c countersServiceImpl) Get(in *counters.CountRequest, out *counters.CountResponse) error {
  out.SetCount(counts[in.GetName()])
  log.Println(fmt.Sprintf("Getting %s: %d", in.GetName(), counts[in.GetName()]))
  return nil
}

func main() {
  var backend = countersServiceImpl{}

  log.Println("Starting service on port 3000")
  counters.ListenAndServeCountersService("0.0.0.0:3000", backend)
}
