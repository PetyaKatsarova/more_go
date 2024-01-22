package main

// import "crypto/x509"
/*
It allows parsing and generating certificates, certificate signing requests, certificate revocation lists, and encoded public and
 private keys. It provides a certificate verifier, complete with a chain builder.The package targets the X.509 technical
  profile defined by the IETF (RFC 2459/3280/5280), and as further restricted by the CA/Browser Forum Baseline Requirements.
   There is minimal support for features outside of these profiles, as the primary goal of the package is to provide compatibility with the publicly trusted TLS certificate ecosystem and its policies and constraints.
On macOS and Windows, certificate verification is handled by system APIs, but the package aims to apply consistent validation
rules across operating systems.
*/

// func generateCert(cn string, caCert *x509.Certificate)

import (
	"fmt"
	"log"
	"sync"
	// "time"
)

// *********** multiple sender test *************************888
// channel c has 0/5 items
// receiver 1 read c(2/5) value "data 1"
// receiver 1 read c(0/5) value "data 3"
// receiver 1 goroutine end (channel closed)
// receiver 2 read c(1/5) value "data 2"
// receiver 2 goroutine end (channel closed)
// finish
func main() {
	c := make(chan string, 5)
	log.Printf("channel c has %d/%d items\n", len(c), cap(c))
	wg := *&sync.WaitGroup{}
	wg.Add(3)
	
	go func() {
		defer wg.Done()
		for s := range c {
			log.Printf("receiver 1 read c(%d/%d) value \"%s\" \n", len(c), cap(c), s)
		}
		fmt.Println("receiver 1 goroutine end (channel closed)")
	}()

	go func() {
		defer wg.Done()
		for s := range c {
			log.Printf("receiver 2 read c(%d/%d) value \"%s\" \n", len(c), cap(c), s)
		}
		fmt.Println("receiver 2 goroutine end (channel closed)")
	}()
	go func() {
		defer wg.Done()
		for s := range c {
			log.Printf("receiver 3 read c(%d/%d) value \"%s\" \n", len(c), cap(c), s)
		}
		fmt.Println("receiver 3 goroutine end (channel closed)")
	}()

	// values are randomly dispatch to the receiver
	c <- "data 1"
	c <- "data 2"
	c <- "data 3"
	c <- "data 4"
	c <- "data 5"
	c <- "data 6"
	c <- "data 77"
	close(c)

	// time.Sleep(100 * time.Millisecond)
	wg.Wait()
	fmt.Println("finished")
}
