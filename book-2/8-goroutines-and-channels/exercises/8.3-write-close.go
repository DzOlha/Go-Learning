package exercises

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

/*
	In netcat3, the interface value conn has the concrete type *net.TCPConn,
	which represents a TCP connection. A TCP connection consists of two halves that may be closed
	independently using its CloseRead and CloseWrite methods. Modify the main goroutine of netcat3
	to close only the write half of the connection so that the program will continue to print
	the final echoes from the reverb1 server even after the standard input has been closed.
*/

var echoPort = "8008"

func RunEchoProgram() {
	go echoServer(echoPort)

	// Give server a moment to start listening before client connects
	time.Sleep(200 * time.Millisecond)

	echoClient(echoPort)
}

// Run clock server that writes echo of the provided input to the client
func echoServer(port string) {
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Clock server started on localhost:" + port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleEchoConn(conn) // handle one connection at a time
	}
}

func handleEchoConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
}

func echoClient(port string) {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatal("connection is not TCP")
	}

	done := make(chan struct{})
	go func() {
		handleCopy(os.Stdout, tcpConn) // copy output into stdout
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	handleCopy(tcpConn, os.Stdin) // copy input from stdin
	tcpConn.CloseWrite()
	<-done // wait for background goroutine to finish
}

func handleCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
