package __wait_group

/*
	Modify the reverb2 server to use a sync.WaitGroup per connection to count the number
	of active echo goroutines. When it falls to zero, close the write half of the TCP connection
	as described in Exercise 8.3. Verify that your modified netcat3 client from that exercise waits
	for the final echoes if multiple concurrent shouts, even after the standard input has been closed.
*/

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var echoPort = "8008"

func RunEchoProgramExercise() {
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
	var wg sync.WaitGroup

	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		s := input.Text()
		go func(text string) {
			defer wg.Done()
			echo(c, text, 1*time.Second)
		}(s)
	}

	go func() {
		wg.Wait()
		if tcp, ok := c.(*net.TCPConn); ok {
			tcp.CloseWrite()
		}
	}()
}

func echoClient(port string) {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	tcp := conn.(*net.TCPConn)
	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		close(done)
		//done <- struct{}{} // signal the main goroutine
	}()
	mustCopyEcho(conn, os.Stdin)
	tcp.CloseWrite()
	<-done // wait for background goroutine to finish
}

func mustCopyEcho(dst io.Writer, src io.Reader) {
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
