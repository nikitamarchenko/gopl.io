/*

Exercise 8.13: Make the chat server disconnect idle clients, such as those that
have sent no messages in the last five minutes. Hint: calling conn.Close() in
another goroutine unblocks active Read calls such as the one done by input.Scan().

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
	//idleDuration = 5 * time.Minute
	idleDuration = 5 * time.Second
)

// !+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	client := client{who, ch}
	ch <- "You are " + who
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
