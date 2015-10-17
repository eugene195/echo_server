package main
//
//import (
//	"fmt"
//	"net"
//	"os"
//	"strconv"
//	"bytes"
//	"flag"
//)

//const (
//	CONN_HOST = ""
//	CONN_PORT = "8080"
//	CONN_TYPE = "tcp"
//)
//
//func main() {
//	wordPtr := flag.String("r", "/", "Root directory")
//	numbPtr := flag.Int("c", 1, "Number of CPU cores")
//	flag.Parse()
//	fmt.Println("word:", *wordPtr)
//	fmt.Println("numb:", *numbPtr)
//
//	// Listen for incoming connections.
//	l, err := net.Listen(CONN_TYPE, ":"+CONN_PORT)
//	if err != nil {
//		fmt.Println("Error listening:", err.Error())
//		os.Exit(1)
//	}
//	// Close the listener when the application closes.
//	defer l.Close()
//	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
//	for {
//		// Listen for an incoming connection.
//		conn, err := l.Accept()
//		if err != nil {
//			fmt.Println("Error accepting: ", err.Error())
//			os.Exit(1)
//		}
//
//		//logs an incoming message
//		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
//
//		// Handle connections in a new goroutine.
//		go handleRequest(conn)
//	}
//}
//
//// Handles incoming requests.
//func handleRequest(conn net.Conn) {
//	// Make a buffer to hold incoming data.
//	buf := make([]byte, 1024)
//	// Read the incoming connection into the buffer.
//	reqLen, err := conn.Read(buf)
//	if err != nil {
//		fmt.Println("Error reading:", err.Error())
//	}
//	// Builds the message.
//	message := "Hi, I received your message! It was "
//	message += strconv.Itoa(reqLen)
//	message += " bytes long and that's what it said: \""
//	n := bytes.Index(buf, []byte{0})
//	message += string(buf[:n-1])
//	message += "\" ! Honestly I have no clue about what to do with your messages, so Bye Bye!\n"
//
//	// Write the message in the connection channel.
//	conn.Write([]byte(message));
//	// Close the connection when you're done with it.
//	conn.Close()
//}

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	NWorkers = flag.Int("n", 4, "The number of workers to start")
	HTTPAddr = flag.String("http", "127.0.0.1:8000", "Address to listen for HTTP requests on")
)

func main() {
	// Parse the command-line flags.
	flag.Parse()

	// Start the dispatcher.
	fmt.Println("Starting the dispatcher")
	StartDispatcher(*NWorkers)

	// Register our collector as an HTTP handler function.
	fmt.Println("Registering the collector")
	http.HandleFunc("/work", Collector)

	// Start the HTTP server!
	fmt.Println("HTTP server listening on", *HTTPAddr)
	if err := http.ListenAndServe(*HTTPAddr, nil); err != nil {
		fmt.Println(err.Error())
	}
}