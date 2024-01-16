package main

import (
	"bytes"
	"log"
	"testing"
)

func Test_Main(t *testing.T) {
	var s bytes.Buffer //capture the standard output and
	// error output of the main function
	log.SetOutput(&s) 
	log.SetFlags(0) // et the flags of the standard logger
	// to 0, which effectively disables any prefix 
	//information (such as timestamps) that would normally be included in log messages. This ensures that the log messages are as clean as possible for testing purposes.
	main()

	if s.String() != "5050 55\n" {
		t.Error(s.String())
	}
}