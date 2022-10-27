package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	m := strings.Fields(ln)[0] //method
	u := strings.Fields(ln)[1] //uri
	fmt.Println("METHOD", m)
	fmt.Println("URI", u)

	if m == "GET" && u == "/" {
		index(conn)
	}
	// if m == "GET" && u == "/about" {
	// 	about(conn)
	// }
	// if m == "GET" && "/contact" {
	// 	about(conn)
	// }
	// if m == "GET" && "/apply" {
	// 	apply(conn)
	// }
	// if m == "POST" && u == "/apply" {
	// 	applyProcess(conn)
	// }
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>INDEX</string><br>
	<a href="/">index</a><br>
	<a href="/">about</a><br>
	<a href="/">contact</a><br>
	<a href="/">apply</a><br>
	</body></html>`
	fmt.Fprint(conn, body)
}

// func about(conn net.Conn) {

// }

// func contact(conn net.Conn) {

// }

// func apply(conn net.Conn) {

// }

// func allpyProcess(conn net.Conn) {

// }
