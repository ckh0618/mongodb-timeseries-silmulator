package sensor

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

type metricType float32
type SubMetric struct {
	Metric01 metricType
	Metric02 metricType
	Metric03 metricType
	Metric04 metricType
	Metric05 metricType
	Metric06 metricType
	Metric07 metricType
	Metric08 metricType
	Metric09 metricType
	Metric10 metricType
}

type NestedDataPoint struct {
	Ts        time.Time
	MetaField map[string]string
	Metric1   SubMetric
	Metric2   SubMetric
	Metric3   SubMetric
	Metric4   SubMetric
	Metric5   SubMetric
}

var MappingFunctionArbitrary = GenArbitraryMetricWithMap
var MappingFunctionNestedObject = GenNestedDocument

func GenNestedDocument(mataFieldMap map[string]string, numOfMetrics int, t time.Time) any {

	subMetric := SubMetric{
		Metric01: metricType(rand.Float32()),
		Metric02: metricType(rand.Float32()),
		Metric03: metricType(rand.Float32()),
		Metric04: metricType(rand.Float32()),
		Metric05: metricType(rand.Float32()),
		Metric06: metricType(rand.Float32()),
		Metric07: metricType(rand.Float32()),
		Metric08: metricType(rand.Float32()),
		Metric09: metricType(rand.Float32()),
		Metric10: metricType(rand.Float32()),
	}
	d := NestedDataPoint{
		Ts:        t.Truncate(time.Second),
		MetaField: mataFieldMap,
		Metric1:   subMetric,
		Metric2:   subMetric,
		Metric3:   subMetric,
		Metric4:   subMetric,
		Metric5:   subMetric,
	}

	return &d
}

func GenArbitraryMetricWithMap(mataFieldMap map[string]string, numOfMetrics int, t time.Time) any {

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
