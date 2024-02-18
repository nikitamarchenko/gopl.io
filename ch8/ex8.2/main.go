/*
Exercise 8.2: Implement a concurrent File Transfer Protocol (FTP) server. The
server should interpret commands from each client such as cd to change
directory, ls to list a directory, get to send the contents of a file, and
close to close the connection. You can use the standard ftp command as the
client, or write your own.
*/

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	r200ok                = "200 Command okay.\r\n"
	r220hello             = "220 Hello\r\n"
	r230loginOk           = "230 User logged in, proceed.\r\n"
	r502notImpl           = "502 Command not implemented.\r\n"
	r522notSup            = "522 Not supported (959)\r\n"
	r211Feat              = "211-Extensions supported:\r\n211 End\r\n"
	r504notImplParam      = "504 Command not implemented for that parameter.\r\n"
	r501errorArg          = "501 Syntax error in parameters or arguments.\r\n"
	r257pwd               = "257 \"%s\" is the current working directory.\r\n"
	r550cwdFail           = "550 Current working directory not changed - %s.\r\n"
	r250cwdSuccess        = "250 \"%s\" is the current working directory.\r\n"
	r150accept            = "150 Accepted data connection\r\n"
	r226CloseConList      = "226 %d matches total\r\n"
	r425ConOpenErr        = "425 Can't open data connection.\r\n"
	r425ConClosedTA       = "426 Connection closed; transfer aborted.\r\n"
	r450ReqActionNotTaken = "450 Requested file action not taken.\r\n"
	r451ReqActionAborted  = "451 Requested action aborted: local error in processing.\r\n"
	r250FileOk            = "250 Requested file action okay, completed.\r\n"
	r221Exit              = "221 Service closing control connection.\r\n"
)

var basicCommands = map[string]string{
	"TYPE A":   r200ok,
	"TYPE A N": r200ok,
	"QUIT":     r221Exit,
	"FEAT":     r211Feat,
	"MODE S":   r200ok,
	"MODE B":   r504notImplParam,
	"MODE C":   r504notImplParam,
	"NOOP":     r200ok,
}

func split2(s string) (string, string, error) {
	ch := strings.Split(s, " ")
	if len(ch) != 2 {
		return "", "", fmt.Errorf("split2 error %s", s)
	}
	return ch[0], ch[1], nil
}

func getCommand(s string) string {
	ch := strings.Split(s, " ")
	return ch[0]
}

func getParam1(s string) string {
	i := strings.Index(s, " ")
	if i > 0 && i+1 < len(s) {
		return s[i+1:]
	}
	return ""
}

func handShake(c net.Conn) error {
	io.WriteString(c, r220hello)
	scanner := bufio.NewScanner(c)
	scanner.Scan()
	log.Println(">", scanner.Text())
	if err := scanner.Err(); err != nil {
		return errors.Join(fmt.Errorf("handShake scanner error"), err)
	}
	command, _, err := split2(scanner.Text())

	if err != nil {
		return errors.Join(fmt.Errorf("handShake error"), err)
	}

	if command != "USER" {
		return fmt.Errorf("handShake error want=USER get=%s", command)
	}
	io.WriteString(c, r230loginOk)
	return nil
}

func getAppDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex) + "/"
}

func scanRead(con net.Conn, c chan string) {
	scanner := bufio.NewScanner(con)
	defer close(c)
	for {
		if r := scanner.Scan(); !r {
			log.Printf("scanRead connection lost\n")
			return
		}
		if err := scanner.Err(); err != nil {
			log.Printf("scanRead scanner error %v\n", err)
			return
		}
		c <- scanner.Text()
	}
}

func handleConn(c net.Conn) {
	defer log.Printf("handleConn exit")
	defer c.Close()
	var err error
	if err = handShake(c); err != nil {
		log.Printf("error %v\n", err)
		return
	}
	appDir := getAppDir()
	cwd := appDir
	pwd := appDir[len(appDir)-1:]
	cAddress := ""

	retChan := make(chan string)
	defer close(retChan)
	scanChan := make(chan string)
	go scanRead(c, scanChan)
	for {
		var cl string
		select {
		case m, ok := <-retChan:
			if ok {
				io.WriteString(c, m)
				continue
			}
		case m, ok := <-scanChan:
			if ok {
				cl = m
			} else {
				return
			}
		default:
			continue
		}

		log.Println(">", cl)
		if resp, ok := basicCommands[cl]; ok {
			io.WriteString(c, resp)
			continue
		}
		switch cl {
		case "PWD":
			io.WriteString(c, fmt.Sprintf(r257pwd, pwd))
			continue
		case "LIST":
			listHandler(c, cAddress, cwd)
			continue
		}

		param1 := getParam1(cl)
		log.Println(">>", param1)
		switch getCommand(cl) {
		case "EPRT":
			io.WriteString(c, r522notSup)
		case "TYPE":
			io.WriteString(c, r501errorArg)
		case "PORT":
			cAddress, err = parsePort(param1)
			if err != nil {
				io.WriteString(c, r501errorArg)
			} else {
				log.Printf("user host:port %s\n", cAddress)
				io.WriteString(c, r200ok)
			}
		case "CWD":
			if cwdN, pwdN, err := cwdHandler(c, param1, appDir, cwd); err == nil {
				cwd, pwd = cwdN, pwdN
			}
		case "LIST":
			listHandlerWithParam(c, cAddress, appDir, cwd, param1)
		case "RETR":
			go retrHandler(retChan, cAddress, appDir, cwd, param1)
		default:
			io.WriteString(c, r502notImpl)
		}
	}
}

func retrHandler(c chan string, host, appDir, cwd, p string) {
	filename, err := fileCalc(p, appDir, cwd)
	if err != nil {
		c <- r451ReqActionAborted
		return
	}
	tc, err := net.Dial("tcp", host)
	if err != nil {
		log.Printf("listHandler: dial error %s", err)
		c <- r425ConOpenErr
		return
	}
	defer tc.Close()
	c <- r150accept
	log.Printf("read file %s\n", filename)
	// not optimal will cause problems on big files
	f, err := os.ReadFile(filename)
	log.Printf("end read file %s\n", filename)
	if err != nil {
		c <- r451ReqActionAborted
		return
	}
	_, err = tc.Write(f)
	if err != nil {
		c <- r425ConClosedTA
	} else {
		c <- r250FileOk
	}
	log.Printf("end of retr\n")
}

func listHandlerWithParam(c net.Conn, host, appDir, cwd, p string) {
	cwd, _, err := cwdCalc(p, appDir, cwd)
	if err != nil {
		io.WriteString(c, r451ReqActionAborted)
		return
	}
	listHandler(c, host, cwd)
}

func listHandler(c net.Conn, host, p string) {
	tc, err := net.Dial("tcp", host)
	if err != nil {
		log.Printf("listHandler: dial error %s", err)
		io.WriteString(c, r425ConOpenErr)
		return
	}
	defer tc.Close()
	io.WriteString(c, r150accept)
	dirs, err := os.ReadDir(p)
	if err != nil {
		io.WriteString(c, r451ReqActionAborted)
		return
	}
	var b strings.Builder
	for _, d := range dirs {
		b.WriteString(d.Name())
		b.WriteString("\r\n")
	}
	_, err = io.WriteString(tc, b.String())
	if err != nil {
		io.WriteString(c, r425ConClosedTA)
	} else {
		io.WriteString(c, fmt.Sprintf(r226CloseConList, len(dirs)))
	}
}

func parsePort(s string) (r string, err error) {
	const LEN = 6
	ss := strings.Split(s, ",")
	if len(ss) != LEN {
		return "", fmt.Errorf("parsePort: invalid format %s", s)
	}
	is := make([]uint32, LEN)
	for i, f := range ss {
		uf, err := strconv.ParseUint(f, 10, 8)
		if err != nil {
			return "", fmt.Errorf("parsePort: invalid format %s, %s not an 8bit uint", s, f)
		}
		is[i] = uint32(uf)
	}

	// we don't need that, just for fun
	var addr uint32 = is[0]
	for i := 1; i < 4; i++ {
		addr <<= 8
		addr += is[i]
	}
	host := fmt.Sprintf("%d.%d.%d.%d",
		uint8(addr>>(8*3)),
		uint8(addr>>(8*2)),
		uint8(addr>>8),
		uint8(addr))

	var port uint16 = uint16(is[4]<<8) + uint16(is[5])

	return fmt.Sprintf("%s:%d", host, port), nil
}

func cwdHandler(c net.Conn, param, appDir, cwd string) (cwdR, pwd string, err error) {
	cwdR, pwd, err = cwdCalc(param, appDir, cwd)
	if err == nil {
		io.WriteString(c, fmt.Sprintf(r250cwdSuccess, pwd))
	} else {
		io.WriteString(c, fmt.Sprintf(r550cwdFail, err))
		return "", "", err
	}
	log.Printf("cwd=%s pwd=%s", cwdR, pwd)
	return
}

func cwdCalc(param, appDir, cwd string) (cwdR, pwd string, err error) {
	newCwd := ""
	if param == "/" {
		newCwd = appDir
	} else if strings.HasPrefix(param, "/") {
		newCwd = filepath.Join(appDir, param)
	} else {
		newCwd = filepath.Join(cwd, param)
	}
	newCwd, err = filepath.Abs(newCwd)
	if err != nil {
		return
	}
	if _, err = os.Stat(newCwd); errors.Is(err, os.ErrNotExist) {
		return
	}
	newCwd += "/"
	if len(newCwd) < len(appDir) {
		return "", "", fmt.Errorf("access denied")
	}
	cwdR = newCwd
	pwd = newCwd[len(appDir)-1:]
	return
}

func fileCalc(param, appDir, cwd string) (cwdR string, err error) {
	newCwd := ""
	if param == "/" {
		return "", fmt.Errorf("not a file")
	} else if strings.HasPrefix(param, "/") {
		newCwd = filepath.Join(appDir, param)
	} else {
		newCwd = filepath.Join(cwd, param)
	}
	newCwd, err = filepath.Abs(newCwd)
	if err != nil {
		return
	}
	if _, err = os.Stat(newCwd); errors.Is(err, os.ErrNotExist) {
		return
	}
	if len(newCwd) < len(appDir) {
		return "", fmt.Errorf("access denied")
	}
	cwdR = newCwd
	return
}

func main() {
	port := flag.Int("port", 8000, "port number")
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}
