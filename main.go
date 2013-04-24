package main

import(
  "encoding/json"
  "bufio"
  "os"
)

func main() {
  for {
    reader := bufio.NewReader(os.Stdin)
    message, _ := reader.ReadString('\n')
    message = message[0:len(message)-1]
    jsonMap := make(map[string]string)
    jsonMap["body"] = message
    json, _ := json.Marshal(jsonMap)
    os.Stdout.Write(json)
    os.Stdout.Write([]byte("\n"))
  }
}
/*
func passMessage(reader, writer) {
  var message string
  reader.read(&message)
  writer.write(message)
}
*/
