// TCP Server para protocolo http.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer listen.Close()

	// Loop para escuchar conexiones entrantes.
	for {
		conexion, err := listen.Accept()
		if err != nil {
			log.Fatal(err.Error())
			continue
		}

		// goroutine que enviamos a la funcion handle
		go handle(conexion)
	}
}

func handle(conexion net.Conn) {
	defer conexion.Close()

	// Leemos el request.
	request(conexion)

	// Escribimos una respuesta
	response(conexion)
}

func request(connexion net.Conn) {
	i := 0
	scanner := bufio.NewScanner(connexion)
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(linea)
		if i == 0 {
			m := strings.Fields(linea)[0]
			fmt.Println("****METHOD", m)
		}
		if linea == "" {
			break
		}
		i++
	}
}

func response(conexion net.Conn) {
	// cuerpo de la respuesta.
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conexion, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conexion, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conexion, "Content-Type: text/html\r\n")
	fmt.Fprint(conexion, "\r\n")
	fmt.Fprint(conexion, body)
}
