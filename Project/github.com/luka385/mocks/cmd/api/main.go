package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	// Crear un contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Configurar las opciones de cliente de MongoDB (como URI)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Reemplaza con tu URI

	// Conectar a MongoDB pasando el contexto y las opciones de cliente
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Verificar si la conexión fue exitosa
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexión exitosa a MongoDB")
}
