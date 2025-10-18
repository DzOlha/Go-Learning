package examples

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

var clockPort = "8008"

func RunClockProgram() {
	go clockServer(clockPort)

	// Give server a moment to start listening before client connects
	time.Sleep(200 * time.Millisecond)

	clockClient(clockPort)
}

// Run clock server that writes the current time to the client once per second
func clockServer(port string) {
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
		go handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func clockClient(port string) {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
