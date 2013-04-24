package main

import(
  "./encoder"
  "os"
  "fmt"
)

func main() {
  var _ = fmt.Printf

  for {
    succ := encoder.PassMessage(os.Stdin, os.Stdout)
    if !succ {
      break
    }
  }
}

