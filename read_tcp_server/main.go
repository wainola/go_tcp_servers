// Leyendo datos al elevar un servidor TCP
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer listen.Close()

	// loopeamos para escuchar la conexion
	for {
		conexion, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// Goroutine con funcion handle
		go handle(conexion)
	}
}

func handle(conexion net.Conn) {
	scanner := bufio.NewScanner(conexion)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conexion.Close()
}
