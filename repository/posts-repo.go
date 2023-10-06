package repository

import (
	"compartamos-clientes/entity"
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type ClienteRepository interface {
	GuardarCliente(cliente *entity.Cliente) (*entity.Cliente, error)
	ListarClientes() ([]entity.Cliente, error)
	ActualizarCliente(cliente *entity.Cliente) (*entity.Cliente, error)
}

type repo struct{}

func NewClienteRepository() ClienteRepository {
	return &repo{}
}

const (
	projectId      string = "compartamos-delacruz"
	collectionName string = "clientes"
)

func (*repo) GuardarCliente(cliente *entity.Cliente) (*entity.Cliente, error) {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	opt := option.WithCredentialsFile(mydir + "\\compartamos-delacruz-firebase-adminsdk-urw1h-5d87bea9c7.json")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, opt)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"Dni":        cliente.Dni,
		"Nombres":    cliente.Nombres,
		"Apellidos":  cliente.Apellidos,
		"Nacimiento": cliente.Nacimiento,
		"Ciudad":     cliente.Ciudad,
		"Direccion":  cliente.Direccion,
		"Correo":     cliente.Correo,
		"Telefono":   cliente.Telefono,
	})
	if err != nil {
		log.Fatalf("error al agregar un nuevo cliente: %v", err)
		return nil, err
	}
	return cliente, nil
}

func (*repo) ActualizarCliente(cliente *entity.Cliente) (*entity.Cliente, error) {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	opt := option.WithCredentialsFile(mydir + "\\compartamos-delacruz-firebase-adminsdk-urw1h-5d87bea9c7.json")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, opt)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	co := client.Doc("pVlWLodA82izkWKgBGRj")
	_, err = co.Update(ctx, []firestore.Update{
		{Path: "Dni", Value: cliente.Dni},
		{Path: "Nombres", Value: cliente.Nombres},
		{Path: "Apellidos", Value: cliente.Apellidos},
		{Path: "Nacimiento", Value: cliente.Nacimiento},
		{Path: "Ciudad", Value: cliente.Ciudad},
		{Path: "Direccion", Value: cliente.Direccion},
		{Path: "Correo", Value: cliente.Correo},
		{Path: "Telefono", Value: cliente.Telefono},
	})

	if err != nil {
		log.Fatalf("error al agregar un nuevo cliente: %v", err)
		return nil, err
	}
	return cliente, nil
}

func (*repo) ListarClientes() ([]entity.Cliente, error) {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	opt := option.WithCredentialsFile(mydir + "\\compartamos-delacruz-firebase-adminsdk-urw1h-5d87bea9c7.json")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, opt)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	var clientes []entity.Cliente
	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		cliente := entity.Cliente{
			Dni:        doc.Data()["Dni"].(string),
			Nombres:    doc.Data()["Nombres"].(string),
			Apellidos:  doc.Data()["Apellidos"].(string),
			Nacimiento: doc.Data()["Nacimiento"].(string),
			Ciudad:     doc.Data()["Ciudad"].(string),
			Direccion:  doc.Data()["Direccion"].(string),
			Correo:     doc.Data()["Correo"].(string),
			Telefono:   doc.Data()["Telefono"].(int64),
		}
		clientes = append(clientes, cliente)
	}
	return clientes, nil
}
