package server

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
	"sync"
	"time"
)

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		// randomly select 1 character from given charset
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func buildMetaFields(n int) map[string]string {

	r := make(map[string]string)

	for i := 0; i < n; i++ {
		k := fmt.Sprintf("MetaFieldKey-%d", i)
		v := randStr(10)

		r[k] = v
	}

	return r
}

type MetricGroup struct {
	startTime     time.Time
	nMetaFields   int
	nMetricFields int
	nMetrics      int
	nBulk         int

	metrics  []*Metric
	channels []chan *bson.D

	wgProduce  *sync.WaitGroup
	wgConsumer *sync.WaitGroup
}

func NewMetricGroup(startTime time.Time, nMetaFields int, nMetricFields int, nMetrics int, nBulk int) *MetricGroup {
	var metrics []*Metric
	var channels []chan *bson.D

	for i := 0; i < nMetrics; i++ {
		meta := buildMetaFields(nMetaFields)
		c := make(chan *bson.D, MaxMetricBuffer)
		channels = append(channels, c)

		p := NewMetricProducer(meta, nMetricFields, startTime, time.Second, c)
		metrics = append(metrics, p)
	}

	return &MetricGroup{
		startTime:     startTime,
		nMetaFields:   nMetaFields,
		nMetricFields: nMetricFields,
		nMetrics:      nMetrics,
		metrics:       metrics,
		channels:      channels,
		nBulk:         nBulk,
		wgConsumer:    new(sync.WaitGroup),
		wgProduce:     new(sync.WaitGroup),
	}
}

func (g MetricGroup) ProduceData(ctx context.Context, count int) error {

	for i := 0; i < g.nMetrics; i++ {
		g.wgProduce.Add(1)
		go func(n int) {
			_ = g.metrics[n].Produce(ctx, count)
			g.wgProduce.Done()
		}(i)
	}
	g.wgProduce.Wait()
	return nil
}

func (g MetricGroup) SubscribeData(ctx context.Context, f func(ctx2 context.Context, dataChan chan *bson.D, nBulk int)) error {

	for i := 0; i < g.nMetrics; i++ {
		g.wgConsumer.Add(1)
		go func(n int) {
			f(ctx, g.channels[n], g.nBulk)
			g.wgConsumer.Done()
		}(i)
	}
	return nil
}

func (g MetricGroup) Wait() {
	g.wgConsumer.Wait()
}
