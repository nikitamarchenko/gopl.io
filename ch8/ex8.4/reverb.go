/*
Exercise 8.4: Modify the reverb2 server to use a sync.WaitGroup per connection
to count the number of active echo goroutines. When it falls to zero, close the
write half of the TCP connection as described in Exercise 8.3. Verify that your
modified netcat3 client from that exercise waits for the final echoes of
multiple concurrent shouts, even after the standard input has been closed.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(wg *sync.WaitGroup, c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	wg.Done()
}

// !+
func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go echo(&wg, c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	wg.Wait()
	tc, ok := c.(*net.TCPConn)
	if ok {
		tc.CloseWrite()
	} else {
		panic("c not a net.TCPConn")
	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
