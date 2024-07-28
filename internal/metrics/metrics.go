package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	issueLabel   = "issue"
	handlerLabel = "handler"
)

var (
	notifiedIssueOrder = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "pvz_issued_orders",
		Help: "total number of issued orders",
	}, []string{
		issueLabel,
	})
)

func AddNotifiedPositionsByContactTotal(cnt int, contact string) {
	notifiedIssueOrder.With(prometheus.Labels{
		issueLabel: contact,
	}).Add(float64(cnt))
}

func IncIssueByHandler(handler string) {
	notifiedIssueOrder.With(prometheus.Labels{
		handlerLabel: handler,
	}).Inc()
}
