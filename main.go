package main

import(
  "fmt"
  "encoding/json"
  "bufio"
  "os"
)

func main() {
  for {
    reader := bufio.NewReader(os.Stdin)
    message, _ := reader.ReadString('\n')
    jsonMap := make(map[string]string)
    jsonMap["body"] = message
    json, _ := json.Marshal(jsonMap)
    fmt.Printf("%s\n", json)
  }
}
/*
func passMessage(reader, writer) {
  var message string
  reader.read(&message)
  writer.write(message)
}
*/
