package encoder

import (
  "encoding/json"
  "bufio"
  "io"
)

func PassMessage(reader io.Reader, writer io.Writer) bool {
  bufReader := bufio.NewReader(reader)
  message, err := bufReader.ReadString('\n')

  if err != nil && err.Error() == "EOF" {
    return false
  }

  message = message[0:len(message)-1]

  jsonMap := make(map[string]string)
  jsonMap["body"] = message
  json, _ := json.Marshal(jsonMap)

  writer.Write(json)
  writer.Write([]byte("\n"))
  return true
}
