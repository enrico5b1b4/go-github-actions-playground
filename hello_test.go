package play_test

import (
	"testing"
  
  play "github.com/enrico5b1b4/go-github-actions-playground"
)

func TestHello(t *testing.T) {
  hello := play.Hello("Enrico")
  if hello != "Hello Enrico" {
    t.Fail()
  }
}
