package main

import (
	"bytes"
	"log"
	"testing"
)

func Test_Main(t *testing.T) {
	var s bytes.Buffer
	log.SetOutput(&s)
	log.SetFlags(0)
	main()

	if s.String() != "Thanks for Hello John\nHello Pip\n" {
		t.Error(s.String())
	}
}