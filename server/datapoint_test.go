package server

import (
	"fmt"
	"testing"
	"time"
)

func TestGenBsonDocument(t *testing.T) {
	d := GenBsonDocument("ServiceName1", "type1", "child1", 4)
	fmt.Println(d)

}

func TestGenArbitraryMetric(t *testing.T) {
	d := GenArbitraryMetric("av", 2, 4, time.Now())
	fmt.Println(d)
}

func TestGenArbitraryMetricWithMap(t *testing.T) {

	m := map[string]string{
		"resourceId":   "serverId_0",
		"definitionId": "server.Filesystem_usage",
	}

	d := GenArbitraryMetricWithMap(m, 4, time.Now())
	fmt.Println(d)
}
