/*

Exercise 8.14: Change the chat server’s network protocol so that each client
provides its name on entering. Use that name instead of the network address
when prefixing each message with its sender’s identity.

*/

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// !+broadcaster
type client struct {
	name string
	ch   chan<- string // an outgoing message channel
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.ch <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			cli.ch <- "Users:"
			for c := range clients {
				cli.ch <- fmt.Sprintf("\t%s", c.name)
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

//!-broadcaster

const (
	idleDuration = 5 * time.Minute
	//idleDuration = 5 * time.Second
)

func identHandler(conn net.Conn) (name string, err error) {

	done := make(chan struct{})
	timeout := time.NewTimer(idleDuration)

	go func() {
		defer func() {
			done <- struct{}{}
		}()
		_, err = fmt.Fprint(conn, "Your name: ")
		if err != nil {
			return
		}
		_, err = fmt.Fscanln(conn, &name)
		if err != nil {
			return
		}
		if len(name) == 0 {
			err = fmt.Errorf("invalid name")
			return
		}
	}()

	select {
	case <-done:
		timeout.Stop()
	case <-timeout.C:
		conn.Close()
		err = fmt.Errorf("timeout")
	}
	return
}

// !+handleConn
func handleConn(conn net.Conn) {
	who, err := identHandler(conn)
	if err != nil {
		log.Printf("%s identification error %s", conn.RemoteAddr().String(), err)
		return
	}
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)
	ch <- "You are " + who
	client := client{who, ch}
	messages <- who + " has arrived"
	entering <- client
	timeout := time.NewTimer(idleDuration)
	end := make(chan struct{})
	endInvoke := func() {
		end <- struct{}{}
	}
	go func() {
		defer endInvoke()
		input := bufio.NewScanner(conn)
		// NOTE: ignoring potential errors from input.Err()
		for input.Scan() {
			if !timeout.Stop() {
				<-timeout.C
			}
			timeout.Reset(idleDuration)
			messages <- who + ": " + input.Text()
		}
	}()

	select {
	case <-timeout.C:
		conn.Close()
		<-end
		break
	case <-end:
		break
	}

	leaving <- client
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

// !+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
