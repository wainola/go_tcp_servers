// MUX que maneja distintas rutas enviadas al servidor.
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
	// Cerramos cuando sea necesario.
	defer listen.Close()

	// Loop del servidor.
	for {
		conexion, err := listen.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}

		// goroutine que envia a la funcion handle
		go handle(conexion)
	}
}

func handle(conexion net.Conn) {
	// defer close y funcion request
	defer conexion.Close()
	request(conexion)
}

func request(conexion net.Conn) {
	// Contador que nos sire para obtener los primeros cambos del http request
	i := 0
	// Retorna un pointer a un escanner que podemos usar.
	scanner := bufio.NewScanner(conexion)
	// suerte de iterador sobre el valor scanner
	for scanner.Scan() {
		linea := scanner.Text()
		if i == 0 {
			fmt.Println(linea)
			// Obtenemos la url y el metodo
			method := strings.Fields(linea)[0]
			fmt.Println(method)
		}
		// fin del request romper bucle
		if linea == "" {
			break
		}
		i++
	}
}
