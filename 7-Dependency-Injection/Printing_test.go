package Dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGree(t *testing.T) {
  buffer := bytes.Buffer{}
  Greet(&buffer, "Ravi")

  got := buffer.String()
  want := "Hello, Ravi"

  if got != want {
    t.Errorf("go %q want %q", got, want)
  }
}
