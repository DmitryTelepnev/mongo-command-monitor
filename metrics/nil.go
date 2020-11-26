package metrics

import "github.com/DmitryTelepnev/mongo-command-monitor/event"

type nil struct {
}

func NewNilMetrics() *nil {
	return &nil{}
}

func (m *nil) CollectStartedQuery(query *event.Query)                 {}
func (m *nil) CollectFinishedQuery(query *event.Query, status string) {}
