package main

import (
	"persistence"
	"time"
)

// Feature ...
type Feature struct {
	ID            int       `json:"id"`          // optional
	DisplayName   string    `json:"displayName"` // optional
	TechnicalName string    `json:"technicalName"`
	ExpiresOn     time.Time `json:"expiresOn"`   // optional
	Description   string    `json:"description"` // optional
	Inverted      bool      `json:"inverted"`
	CustomerIds   []int     `json:"customerIds"`
}

// GetFeatures ...
func GetFeatures() []byte {
	return persistence.GetFeatures()
}

// GetFeature ...
func GetFeature(id int) []byte {
	return persistence.GetFeature(id)
}

// CreateNewFeature ...
func CreateNewFeature(values []byte) bool {
	return persistence.CreateFeature(values)
}

// UpdateFreature ...
func UpdateFreature(values []byte) bool {
	return persistence.UpdateFeature(values)
}

// ToggleFeature ...
func ToggleFeature(id int) bool {
	return persistence.ToggleFeature(id)
}

// ArchiveFeature ...
func ArchiveFeature(id int) bool {
	return persistence.ArchiveFeature(id)
}
