package server

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const MaxMetricBuffer = 60

type Metric struct {
	metaField       map[string]string
	numberOfMetrics int
	startTime       time.Time
	duration        time.Duration
	c               chan *bson.D
}

func NewMetricProducer(metaField map[string]string, numberOfMetrics int, startTime time.Time, duration time.Duration, c chan *bson.D) *Metric {
	return &Metric{metaField: metaField, numberOfMetrics: numberOfMetrics, startTime: startTime, duration: duration, c: c}
}

func (m Metric) Produce(ctx context.Context, cnt int) error {

	t := m.startTime
	for i := 0; i < cnt; i++ {
		insTime := t.Add(time.Duration(i) * m.duration)
		d := GenArbitraryMetricWithMap(m.metaField, m.numberOfMetrics, insTime)
		m.c <- d
		d = nil
		time.Sleep(m.duration)
	}

	close(m.c)
	return nil
}
