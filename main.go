package main

import "fmt"
import "encoding/json"

func main() {
  for {
    var message string;
    jsonMap := make(map[string]string)

    fmt.Scanf("%s", &message)
    jsonMap["body"] = message
    json, _ := json.Marshal(jsonMap)
    fmt.Printf("%s\n", json)
  }
}

