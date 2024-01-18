// tutorial from: https://notes.shichao.io/gopl/ch8/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// Most echo servers merely write whatever they read, which can be done with this trivial version of handleConn:

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	time.Sleep(delay)
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	 // NOTE: ignoring potential errors from input.Err()
	 c.Close()
}

func mustCopy(dst io.Writer, src io.Reader) {
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		 log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

/*
net.Listen is used for setting up servers to accept incoming connections, while net.Dial is used for clients to initiate
 connections to remote servers. 
*/