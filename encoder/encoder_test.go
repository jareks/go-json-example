package encoder

import (
  "testing"
  "bytes"
)

func Test_PassMessage(t *testing.T) {
  input := bytes.NewBufferString("test elo\n")
  output := bytes.NewBufferString("")
  PassMessage(input, output)
  result := output.String()
  if result == "{\"body\":\"test elo\"}\n" {
    t.Log("OK")
  } else {
    t.Log(result)
    t.Fail()
  }
}

