package sensor

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	uri := "mongodb://localhost:27017/test"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	assert.NoError(t, err)
	mongoHandler := NewMongoHandler(client, "ts", "ts")
	group := NewMetricGroup(time.Now(), 3, 1, 10, 10, true)
	group.SubscribeData(context.Background(), mongoHandler.DoInsert)
	group.ProduceData(context.Background(), 100)

	group.Wait()
}
