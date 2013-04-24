package main

import "fmt"
import "encoding/json"

func main() {
  var message string;
  fmt.Scanf("%s", &message)
  json, _ := json.Marshal(message)
  fmt.Printf("%s", json)
}

