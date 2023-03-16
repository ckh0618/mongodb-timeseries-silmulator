package sensor

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoHandler struct {
	client *mongo.Client
	coll   *mongo.Collection

	database   string
	collection string
}

func NewMongoHandler(client *mongo.Client, database string, collection string) *MongoHandler {

	coll := client.Database(database).Collection(collection)
	return &MongoHandler{client: client, database: database, coll: coll, collection: collection}
}

func (m MongoHandler) DoInsert(ctx context.Context, chanData chan any, nBulk int) {

	i := 0
	var arrData []any

	for data := range chanData {
		arrData = append(arrData, data)
		i++
		if i >= nBulk {
			_, err := m.coll.InsertMany(ctx, arrData)
			if err != nil {
				fmt.Println(err)
			}

			i = 0
			arrData = arrData[0:0]
			data = nil
			continue
		}
		data = nil
	}

}
