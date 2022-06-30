package model

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	TotalCountSend = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "count_of_send_messages",
			Help: "Total count of send messages",
		},
		[]string{"notion_send"},
	)

	TotalCountReq = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "count_of_req",
			Help: "Total count of send messages",
		},
		[]string{"notion_req"},
	)
)
