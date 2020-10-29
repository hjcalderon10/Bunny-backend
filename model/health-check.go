package model

import "time"

type HealthCheck struct {
	Name    string    `json:"name"`
	Version string    `json:"version"`
	Date    time.Time `json:"date"`
}
