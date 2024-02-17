/*
Exercise 8.1: Modify clock2 to accept a port number, and write a program,
clockwall, that acts as a client of several clock servers at once, reading the
times from each one and displaying the results in a table, akin to the wall of
clocks seen in some business offices. If you have access to geographically
distributed computers, run instances remotely; otherwise run local instances
on different ports with fake time zones.


$ TZ=US/Eastern    ./clock2 -port 8010 &
$ TZ=Asia/Tokyo    ./clock2 -port 8020 &
$ TZ=Europe/London ./clock2 -port 8030 &
$ clockwall NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030

*/

package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Clock struct {
	label string
	conn  net.Conn
}

func fatalPrintUsage() {
	fmt.Printf("usage: %s ClockName=host:port\n", os.Args[0])
	os.Exit(1)
}

func (c Clock) Close() {
	fmt.Printf("closing connection")
	c.conn.Close()
}

func main() {

	if len(os.Args) == 1 {
		fatalPrintUsage()
	}

	clocks := make([]Clock, 0, len(os.Args)-1)

	for _, w := range os.Args[1:] {
		if !strings.Contains(w, "=") {
			fatalPrintUsage()
		}
		ss := strings.Split(w, "=")
		conn, err := net.Dial("tcp", ss[1])
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		clocks = append(clocks, Clock{ss[0], conn})
	}
	buf := []byte(time.Now().Format("15:04:05\n"))

	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	for len(signal_chan) == 0 {
		for _, c := range clocks {
			n, err := c.conn.Read(buf)
			if err != nil {
				log.Printf("%s error %v\n", c.label, err)
				continue
			}
			if n != len(buf) {
				log.Printf("%s invalid packet size: get %d want %d\n", c.label, n, len(buf))
			}
			fmt.Printf("%s %s ", c.label, string(buf[:len(buf)-1]))
		}
		fmt.Println()
	}
	fmt.Println("Exiting...")
}
