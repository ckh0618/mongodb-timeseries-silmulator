package sensor

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestMetricProducer_Produce(t *testing.T) {

	c := make(chan any, MaxMetricBuffer)

	go func() {

		for {
			d := <-c
			fmt.Println(d)
		}
	}()

	m := map[string]string{
		"a": "a",
		"b": "b",
	}

	go func() {
		producer := NewMetricProducer(m, 3, time.Now(), time.Second, c, MappingFunctionNestedObject)
		producer.Produce(context.Background(), 1000)
	}()

	producer := NewMetricProducer(m, 6, time.Now(), time.Second, c, MappingFunctionNestedObject)
	producer.Produce(context.Background(), 1000)

}
