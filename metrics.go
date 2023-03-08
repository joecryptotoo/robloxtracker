package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	UserPresenceType prometheus.Gauge
}

func robloxMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		UserPresenceType: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "UserPresenceType",
			Help: "Offline, Online, InGame, InStudio, Unknown",
		}),
	}

	reg.MustRegister(m.UserPresenceType)

	return m
}

func updateMetrics(user User) {

	user.Metrics.UserPresenceType.Set(float64(user.Presence.UserPresenceType))

}
