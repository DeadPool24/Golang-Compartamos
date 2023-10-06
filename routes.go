package main

import (
	"encoding/json"
	"net/http"

	"compartamos-clientes/entity"
	"compartamos-clientes/repository"
)

var (
	repo repository.ClienteRepository = repository.NewClienteRepository()
)

func listerClientes(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	clientes, err := repo.ListarClientes()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error al listar clientes"}`))
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(clientes)
}

func guardarCliente(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var cliente entity.Cliente
	err := json.NewDecoder(request.Body).Decode(&cliente)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error al guardar clientes"}`))
		return
	}
	repo.GuardarCliente(&cliente)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(cliente)
}

func actualizarCliente(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var cliente entity.Cliente
	err := json.NewDecoder(request.Body).Decode(&cliente)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error al guardar clientes"}`))
		return
	}
	repo.ActualizarCliente(&cliente)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(cliente)
}
