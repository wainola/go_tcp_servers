// MUX que maneja distintas rutas enviadas al servidor.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"text/template"
)

// Inicializacion de templates.
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.html"))
}

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
			// Obtenemos la url y el metodo
			method := strings.Fields(linea)[0]
			url := strings.Fields(linea)[1]
			fmt.Println("#### METHOD: ", method)
			fmt.Println("#### URL: ", url)
			mux(url, method, conexion)
		}
		// fin del request romper bucle. El objetivo es que, enviada la data, hacer el retorno del request como se haria en una rest api normal
		if linea == "" {
			break
		}
		i++
	}
}

func mux(url string, method string, conexion net.Conn) {
	// Aca se lleva a cabo la logica de programacion.
	if method == "GET" && url == "/" {
		fmt.Println("Index request")
		t := tpl.ExecuteTemplate(os.Stdout, "index.html", nil)
		if t != nil {
			log.Fatal(t)
		}
	}
}
