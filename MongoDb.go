package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoDbStorage() (*MongoStorage, error) {
	ms := MongoStorage{}

	// Si usas un usuario y contraseña, sería "mongodb://user:password@localhost:27017"
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 2. Conectarse a MongoDB
	// Utiliza un contexto con timeout para evitar esperas infinitas.
	ms.ctx, ms.cancel = context.WithTimeout(context.Background(), 10*time.Second)

	//defer ms.cancel() // Asegura que el contexto se cancele cuando la función main termine
	client, err := mongo.Connect(ms.ctx, clientOptions)
	if err != nil {
		return nil, errors.New("Error al conectar a MongoDB: " + err.Error())
	}

	// Asegúrate de cerrar la conexión cuando la función main termine
	/*defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("Error al desconectar de MongoDB: %v", err)
		}
		fmt.Println("Desconectado de MongoDB.")
	}()*/

	// 3. Hacer un "ping" para verificar la conexión
	// readpref.Primary significa que se enviará la operación al nodo primario del conjunto de réplicas
	if err = client.Ping(ms.ctx, readpref.Primary()); err != nil {
		log.Fatalf("Error al hacer ping a MongoDB: %v", err)
	}
	fmt.Println("Conectado exitosamente a MongoDB!")

	// 4. Seleccionar la base de datos y la colección
	ms.client = client
	database := client.Database("Gimnasio")      // Reemplaza "mydatabase" con el nombre de tu base de datos
	collection := database.Collection("alumnos") // Reemplaza "mycollection" con el nombre de tu colección
	ms.coleccion = collection

	// 5. Insertar un documento
	/*fmt.Println("\nInsertando un documento...")
	result, err := collection.InsertOne(ctx, bson.D{
		{Key: "name", Value: "GoLang User"},
		{Key: "age", Value: 30},
		{Key: "email", Value: "go@example.com"},
	})
	if err != nil {
		log.Fatalf("Error al insertar documento: %v", err)
	}
	fmt.Printf("Documento insertado con ID: %v\n", result.InsertedID)

	fmt.Println("\nInsertando un documento...")
	result, err = collection.InsertOne(ctx, bson.D{
		{Key: "name", Value: "Alumno"},
		{Key: "age", Value: 56},
		{Key: "email", Value: "miguelponce1@gmail.com"},
	})
	if err != nil {
		log.Fatalf("Error al insertar documento: %v", err)
	}
	fmt.Printf("Documento insertado con ID: %v\n", result.InsertedID)*/

	// 6. Buscar un documento
	/*fmt.Println("\nBuscando un documento...")
	var user struct {
		Name  string `bson:"name"`
		Age   int    `bson:"age"`
		Email string `bson:"email"`
	}

	filter := bson.D{{Key: "name", Value: "GoLang User"}}
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No se encontró ningún documento con ese nombre.")
		} else {
			log.Fatalf("Error al buscar documento: %v", err)
		}
	} else {
		fmt.Printf("Documento encontrado: %+v\n", user)
	}*/

	// 7. Actualizar un documento (opcional)
	/*fmt.Println("\nActualizando un documento...")
	updateResult, err := collection.UpdateOne(
		ctx,
		bson.D{{Key: "name", Value: "GoLang User"}},
		bson.D{{Key: "$set", Value: bson.D{{Key: "age", Value: 31}}}},
	)
	if err != nil {
		log.Fatalf("Error al actualizar documento: %v", err)
	}
	fmt.Printf("Documentos actualizados: %v\n", updateResult.ModifiedCount)*/

	/*// 8. Eliminar un documento (opcional)
	fmt.Println("\nEliminando un documento...")
	deleteResult, err := collection.DeleteOne(ctx, bson.D{{Key: "name", Value: "GoLang User"}})
	if err != nil {
		log.Fatalf("Error al eliminar documento: %v", err)
	}
	fmt.Printf("Documentos eliminados: %v\n", deleteResult.DeletedCount)*/

	return &ms, nil
}

type MongoStorage struct {
	coleccion *mongo.Collection
	client    *mongo.Client
	ctx       context.Context
	cancel    context.CancelFunc
}

func (ms *MongoStorage) SetAlumno() error {
	// 5. Insertar un documento
	fmt.Println("\nInsertando un documento...")

	result, err := ms.coleccion.InsertOne(ms.ctx, bson.D{
		{Key: "name", Value: "GoLang User"},
		{Key: "age", Value: 30},
		{Key: "email", Value: "go@example.com"},
	})
	if err != nil {
		log.Fatalf("Error al insertar documento: %v", err)
	}
	fmt.Printf("Documento insertado con ID: %v\n", result.InsertedID)

	fmt.Println("\nInsertando un documento...")
	result, err = ms.coleccion.InsertOne(ms.ctx, bson.D{
		{Key: "name", Value: "Alumno"},
		{Key: "age", Value: 56},
		{Key: "email", Value: "miguelponce1@gmail.com"},
	})
	if err != nil {
		log.Fatalf("Error al insertar documento: %v", err)
	}
	fmt.Printf("Documento insertado con ID: %v\n", result.InsertedID)
	return nil
}

func (ms *MongoStorage) GetAlumno() error {
	// 6. Buscar un documento
	fmt.Println("\nBuscando un documento...")
	var user struct {
		Name  string `bson:"name"`
		Age   int    `bson:"age"`
		Email string `bson:"email"`
	}

	filter := bson.D{{Key: "name", Value: "GoLang User"}}
	err := ms.coleccion.FindOne(ms.ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No se encontró ningún documento con ese nombre.")
		} else {
			log.Fatalf("Error al buscar documento: %v", err)
		}
	} else {
		fmt.Printf("Documento encontrado: %+v\n", user)
	}

	return nil
}

func (ms *MongoStorage) UpdateAlumno() error {
	// 7. Actualizar un documento (opcional)
	fmt.Println("\nActualizando un documento...")
	updateResult, err := ms.coleccion.UpdateOne(
		ms.ctx,
		bson.D{{Key: "name", Value: "GoLang User"}},
		bson.D{{Key: "$set", Value: bson.D{{Key: "age", Value: 31}}}},
	)
	if err != nil {
		log.Fatalf("Error al actualizar documento: %v", err)
	}
	fmt.Printf("Documentos actualizados: %v\n", updateResult.ModifiedCount)

	return nil
}

func (ms *MongoStorage) DeleteAlumno() error {
	// 8. Eliminar un documento (opcional)
	fmt.Println("\nEliminando un documento...")
	deleteResult, err := ms.coleccion.DeleteOne(ms.ctx, bson.D{{Key: "name", Value: "GoLang User"}})
	if err != nil {
		log.Fatalf("Error al eliminar documento: %v", err)
	}
	fmt.Printf("Documentos eliminados: %v\n", deleteResult.DeletedCount)
	return nil
}

func (ms *MongoStorage) SetProfesor() error {
	return nil
}

func (ms *MongoStorage) SetTurno() error {
	return nil
}
func (ms *MongoStorage) SetHorario() error {
	return nil
}
func (ms *MongoStorage) GetAlumnto() {
	return
}
func (ms *MongoStorage) GetProfesor() {
	return
}

func (ms *MongoStorage) Desconectar() {
	return
}

func NewMongoDb() *MongoStorage {
	ms := MongoStorage{}
	return &ms
}
