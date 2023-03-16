package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
	"timeseries_test/sensor"
)

func main() {

	collectionName := flag.String("collection", "timeseries", "Collection to insert ")
	databaseName := flag.String("database", "timeseries", "database to insert ")
	nMetaField := flag.Int("metafield-count", 3, "metafield count")
	nMetricField := flag.Int("metric-count", 5, "metric field count")
	nSensors := flag.Int("sensor", 10, "# of sensors")
	nCount := flag.Int("iteration", 1000, "number of test set ")
	nBulk := flag.Int("bulk", 10, "number of batch count")
	nNested := flag.Bool("isNested", false, "true if testing with nested objects")

	flag.Parse()

	startTime := time.Now()
	uri, ok := os.LookupEnv("MONGODB_URI")
	if !ok {
		fmt.Println("MONGODB_URI is not set !")
		return
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("Could not connect ... URI : [%s]\n", uri)
		return
	}

	err = buildTimeSeriesCollection(client, *databaseName, *collectionName)

	if err != nil {
		fmt.Println(err)
	}

	mongoHandler := sensor.NewMongoHandler(client, *databaseName, *collectionName)

	group := sensor.NewMetricGroup(time.Now(), *nMetaField, *nMetricField, *nSensors, *nBulk, *nNested)
	group.SubscribeData(context.Background(), mongoHandler.DoInsert)
	group.ProduceData(context.Background(), *nCount)

	group.Wait()

	endTime := time.Now()

	resultCollection := client.Database(*databaseName).Collection("result")

	testResult := TestResult{
		StartTime:        startTime,
		EndTime:          endTime,
		MetaFieldCount:   *nMetaField,
		MetricFieldCount: *nMetricField,
		SensorCount:      *nSensors,
		BulkCount:        *nBulk,
		IterationCount:   *nCount,
	}

	_, err = resultCollection.InsertOne(context.Background(), testResult)
	if err != nil {
		fmt.Printf("Error !!, [%v]", err)
	}

	return
}

type TestResult struct {
	StartTime        time.Time
	EndTime          time.Time
	MetaFieldCount   int
	MetricFieldCount int
	SensorCount      int
	BulkCount        int
	IterationCount   int
}

func buildTimeSeriesCollection(client *mongo.Client, databaseName string, collectionName string) error {

	database := client.Database(databaseName)

	err := database.CreateCollection(context.Background(), collectionName, &options.CreateCollectionOptions{
		TimeSeriesOptions: &options.TimeSeriesOptions{
			TimeField:   "ts",
			MetaField:   aws.String("metadata"),
			Granularity: aws.String("seconds"),
		},
	})

	if err != nil {
		return err
	}

	coll := database.Collection(collectionName)
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{"metadata", 1},
			{"ts", -1},
		},
	}
	_, err = coll.Indexes().CreateOne(context.Background(), indexModel)

	if err != nil {
		return err
	}

	return nil
}
