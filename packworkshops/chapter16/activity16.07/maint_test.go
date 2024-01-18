package main

import (
	"bytes"
	"log"
	"testing"
)

func Test_Main(t *testing.T) {
	var s bytes.Buffer
	log.SetOutput(&s)
	main()

	output	:= s.String()
	ln		:= len(output)
	res		:= output[ln - 5 : ln - 1]

	if res != "5050" {
		t.Errorf("Expected 5050 but received %s", res)
	}
}