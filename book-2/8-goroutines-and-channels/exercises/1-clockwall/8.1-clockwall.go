package exercises

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

/*
	Modify clock2 to accept a port number, and write a program, clockwall,
	that acts as a client of several clock servers at once, reading the times from each one
	and displaying the results in a table, akin to the wall of clocks seen in some business offices.
	if you have access to geographically distributed computers, run instances remotely; otherwise
	run local instances on different ports with fake time zones.
*/

func RunExampleOfClockWall() {

	go func() {
		os.Args = []string{"cmd", "server", "-tz", "US/Eastern", "-port", "8010"}
		RunClockProgram()
	}()

	go func() {
		os.Args = []string{"cmd", "server", "-tz", "Asia/Tokyo", "-port", "8020"}
		RunClockProgram()
	}()

	time.Sleep(500 * time.Millisecond)

	// Run the wall
	os.Args = []string{"cmd", "wall", "NewYork=localhost:8010", "Tokyo=localhost:8020"}
	RunClockProgram()
}
func RunClockProgram() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println(" clock2 server -tz TIMEZONE -port PORT")
		fmt.Println(" clock2 wall Name=Address ...")
	}
	switch os.Args[1] {
	case "server":
		runServer(os.Args[2:])
	case "wall":
		runWall(os.Args[2:])
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}

func runServer(args []string) {
	fs := flag.NewFlagSet("server", flag.ExitOnError)
	tz := fs.String("tz", "UTC", "Timezone name (e.g., US/Eastern)")
	port := fs.String("port", "8000", "Port to listen on")
	fs.Parse(args)

	location, err := time.LoadLocation(*tz)
	if err != nil {
		log.Fatalf("Invalid timezone: %v", err)
	}
	startClockServer(*port, location)
}

func startClockServer(port string, location *time.Location) {
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
		go handleConn(conn, location) // handle one connection at a time
	}
}

func handleConn(c net.Conn, loc *time.Location) {
	defer c.Close()
	for {
		current := time.Now().In(loc).Format("15:04:05")
		_, err := io.WriteString(c, current+"\n")
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func runWall(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: clock2 wall Name=Address ...")
		return
	}

	count := len(args)
	names := make([]string, count)
	times := make([]string, count)

	for i, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			log.Fatalf("Invalid format: %s (expected Name=host:port)", arg)
		}
		name, addr := parts[0], parts[1]
		names[i] = name
		index := i

		go func(addr string, name string) {
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				log.Printf("Error connecting to %s: %v", name, err)
				times[index] = "ERROR"
				return
			}
			defer conn.Close()

			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				times[index] = scanner.Text()
			}
		}(addr, name)
	}

	for {
		time.Sleep(1 * time.Second)

		for i := 0; i < count; i++ {
			fmt.Printf("\r%-10s %s\n", names[i]+":", times[i])
		}
	}
}
