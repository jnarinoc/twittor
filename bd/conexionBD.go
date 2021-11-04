package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoC exporta la conexion*/
var MongoC = conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:admin@cluster.a9qbv.mongodb.net/myFirstDatabase?authSource=admin&replicaSet=atlas-p67hw1-shard-0&w=majority&readPreference=primary&appname=MongoDB%20Compass&retryWrites=true&ssl=true")

/*ConectarBD es la funcion que me permite conectarme a la bd*/
func conectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}
	log.Print("conexion exitosa con la BD")
	return client
}

func ChequeoConnection() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
