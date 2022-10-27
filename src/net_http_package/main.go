package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":23")
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
	request(conn) //read request
	respond(conn) // write response
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0] //request line
			fmt.Println("***METHOD", m)
		}
		if ln == "" {
			break // headers are done
		}
		i++
	}
}

func respond(conn net.Conn) {
	body := `HERE BODY`
	fmt.Fprint(conn, body)
	fmt.Println("wonder? ", conn)
}
