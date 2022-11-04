package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func concurrente(w http.ResponseWriter, nombre []string) {
	for i := 0; i < len(nombre); i++ {
		io.WriteString(w, "El individuo "+nombre[i]+" esta comiendo")
		fmt.Fprintf(w, "\n")
		fmt.Println("El individuo " + nombre[i] + " esta comiendo")
		time.Sleep(time.Millisecond * 2000)
	}
}

func main() {
	//arreglo
	nombre := []string{"Pedro", "Anita", "Carlos"}

	//browser
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(nombre)
		for i := 0; i < len(nombre); i++ {
			io.WriteString(w, nombre[i])
			fmt.Fprintf(w, "\n")
		}
	})
	http.HandleFunc("/agregar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Ingrese nombre")
		var eleccion string
		fmt.Scanln(&eleccion)
		registro := len(nombre)
		nombre = append(nombre, eleccion)
		fmt.Println("Nombre agregado: " + eleccion)
		time.Sleep(8 * time.Second)
		io.WriteString(w, "El dato agregado fue: "+nombre[registro])
	})
	http.HandleFunc("/comer", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < len(nombre); i++ {
			io.WriteString(w, "El individuo "+nombre[i]+" esta comiendo")
			fmt.Fprintf(w, "\n")
			time.Sleep(8 * time.Second)
		}
		go concurrente(w, nombre)
	})

	//consola
	fmt.Println("El servidor escucha en el puerto 8080")
	http.ListenAndServe(":8080", nil)
}

