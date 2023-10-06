package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})
	router.HandleFunc("/clientes", listerClientes).Methods("GET")
	router.HandleFunc("/clientes", guardarCliente).Methods("POST")
	router.HandleFunc("/clientes", actualizarCliente).Methods("PUT")
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
