package main

import (
	"time"

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
			Help: "Seconds spent offline",
		}),
		OnlineTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "OnlineTime",
			Help: "Seconds spent online",
		}),
		InGameTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "InGameTime",
			Help: "Seconds spent in-game",
		}),
		InStudioTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "InStudioTime",
			Help: "Seconds spent in-studio",
		}),
		UnknownTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "UnknownTime",
			Help: "Seconds spent in an unknown state",
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

func updateMetrics(user User) {

	user.Metrics.UserPresenceType.Set(float64(user.LastPresenceType))

	switch user.LastPresenceType {
	case 0:
		user.Metrics.OfflineTime.Set(float64(time.Now().UTC().Sub(user.LastPresenceChange).Seconds()))
		user.Metrics.OnlineTime.Set(float64(0))
		user.Metrics.InGameTime.Set(float64(0))
		user.Metrics.InStudioTime.Set(float64(0))
		user.Metrics.UnknownTime.Set(float64(0))
	case 1:
		user.Metrics.OfflineTime.Set(float64(0))
		user.Metrics.OnlineTime.Set(float64(time.Now().UTC().Sub(user.LastPresenceChange).Seconds()))
		user.Metrics.InGameTime.Set(float64(0))
		user.Metrics.InStudioTime.Set(float64(0))
		user.Metrics.UnknownTime.Set(float64(0))
	case 2:
		user.Metrics.OfflineTime.Set(float64(0))
		user.Metrics.OnlineTime.Set(float64(0))
		user.Metrics.InGameTime.Set(float64(time.Now().UTC().Sub(user.LastPresenceChange).Seconds()))
		user.Metrics.InStudioTime.Set(float64(0))
		user.Metrics.UnknownTime.Set(float64(0))
	case 3:
		user.Metrics.OfflineTime.Set(float64(0))
		user.Metrics.OnlineTime.Set(float64(0))
		user.Metrics.InGameTime.Set(float64(0))
		user.Metrics.InStudioTime.Set(float64(time.Now().UTC().Sub(user.LastPresenceChange).Seconds()))
		user.Metrics.UnknownTime.Set(float64(0))
	default:
		user.Metrics.OfflineTime.Set(float64(0))
		user.Metrics.OnlineTime.Set(float64(0))
		user.Metrics.InGameTime.Set(float64(0))
		user.Metrics.InStudioTime.Set(float64(0))
		user.Metrics.UnknownTime.Set(float64(time.Now().UTC().Sub(user.LastPresenceChange).Seconds()))
	}
}
