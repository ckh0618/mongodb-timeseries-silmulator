package sensor

import (
	"context"
	"time"
)

const MaxMetricBuffer = 60

type BuilderFunc func(mataFieldMap map[string]string, numOfMetrics int, t time.Time) any
type Metric struct {
	metaField       map[string]string
	numberOfMetrics int
	startTime       time.Time
	duration        time.Duration
	c               chan any

	f BuilderFunc
}

func NewMetricProducer(metaField map[string]string, numberOfMetrics int, startTime time.Time, duration time.Duration, c chan any, f BuilderFunc) *Metric {
	return &Metric{metaField: metaField, numberOfMetrics: numberOfMetrics, startTime: startTime, duration: duration, c: c, f: f}
}

func (m Metric) Produce(ctx context.Context, cnt int) error {

	t := m.startTime
	for i := 0; i < cnt; i++ {
		insTime := t.Add(time.Duration(i) * m.duration)
		d := m.f(m.metaField, m.numberOfMetrics, insTime)
		m.c <- d
		d = nil
		time.Sleep(m.duration)
	}

	close(m.c)
	return nil
}
