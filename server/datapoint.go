package server

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
	"time"
)

type DataPoint struct {
	Timestamp time.Time
	MetaField map[string]string
	Metrics   map[string]float32
}

const metaFieldNamePrefix = "MetaField%d"
const metricFieldNamePrefix = "metrics%d"

func GenArbitraryMetricWithMap(mataFieldMap map[string]string, numOfMetrics int, t time.Time) *bson.D {

	d := bson.D{
		{
			Key:   "ts",
			Value: t.Truncate(time.Second),
		},
	}

	metaFields := bson.D{}

	for k, v := range mataFieldMap {
		m := bson.E{
			Key:   k,
			Value: v,
		}
		metaFields = append(metaFields, m)

	}

	d = append(d, bson.E{
		Key:   "metadata",
		Value: metaFields,
	})

	for i := 0; i < numOfMetrics; i++ {
		metricKey := fmt.Sprintf(metricFieldNamePrefix, i)
		metricValue := rand.Float32()

		val := bson.E{Key: metricKey, Value: metricValue}
		d = append(d, val)

	}

	return &d
}

func GenArbitraryMetric(metaFieldValuePrefix string, numOfMetaField, numOfMetrics int, t time.Time) *bson.D {

	d := bson.D{
		{
			Key:   "ts",
			Value: t.Truncate(time.Second),
		},
	}

	metaFields := bson.D{}

	for i := 0; i < numOfMetaField; i++ {
		m := bson.E{
			Key:   fmt.Sprintf(metaFieldNamePrefix, i),
			Value: fmt.Sprintf(metaFieldValuePrefix+"%d", i),
		}
		metaFields = append(metaFields, m)
	}

	d = append(d, bson.E{
		Key:   "metadata",
		Value: metaFields,
	})

	for i := 0; i < numOfMetrics; i++ {
		metricKey := fmt.Sprintf(metricFieldNamePrefix, i)
		metricValue := rand.Float32()

		val := bson.E{Key: metricKey, Value: metricValue}
		d = append(d, val)

	}

	return &d
}

func GenBsonDocument(resourceID string, resourceType string, childName string, numOfMetrics int) *bson.D {

	d := bson.D{
		{"metaField", bson.M{"resourceId": resourceID, "resourceType": resourceType, "childName": childName}},
	}
	d = append(d, bson.E{"timestamp", time.Now().Truncate(time.Second)})
	for i := 0; i < numOfMetrics; i++ {
		metricKey := "metric" + fmt.Sprintf("%d", i)
		metricValue := rand.Float32()

		val := bson.E{Key: metricKey, Value: metricValue}
		d = append(d, val)

	}
	return &d
}
