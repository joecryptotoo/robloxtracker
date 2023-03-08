package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	UserPresenceType prometheus.Gauge
	OfflineTime      prometheus.Gauge
	OnlineTime       prometheus.Gauge
	InGameTime       prometheus.Gauge
	InStudioTime     prometheus.Gauge
	UnknownTime      prometheus.Gauge
}

func RobloxMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		UserPresenceType: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "UserPresenceType",
			Help: "Offline, Online, InGame, InStudio, Unknown",
		}),
		OfflineTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "OfflineTime",
			Help: "Time spent offline",
		}),
		OnlineTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "OnlineTime",
			Help: "Time spent online",
		}),
		InGameTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "InGameTime",
			Help: "Time spent in-game",
		}),
		InStudioTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "InStudioTime",
			Help: "Time spent in-studio",
		}),
		UnknownTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "UnknownTime",
			Help: "Time spent in an unknown state",
		}),
	}

	reg.MustRegister(m.UserPresenceType)
	reg.MustRegister(m.OfflineTime)
	reg.MustRegister(m.OnlineTime)
	reg.MustRegister(m.InGameTime)
	reg.MustRegister(m.InStudioTime)
	reg.MustRegister(m.UnknownTime)
	return m
}
