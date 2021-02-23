package main

import (
	"encoding/json"
	"log"
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

type feature struct {
	ID            int       `json:"featureId"`
	DisplayName   string    `json:"displayName"`
	TechnicalName string    `json:"technicalName"`
	ExpiresOn     time.Time `json:"expiresOn"`
	Description   string    `json:"description"`
	Inverted      bool      `json:"inverted"`
	Active        bool      `json:"active"`
}

type features struct {
	Features []feature `json:"features"`
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
func UpdateFreature(id int, values []byte) bool {
	var feat feature
	json.Unmarshal(GetFeatures(), &feat)
	feat.ID = id

	var jsonData []byte
	jsonData, err := json.Marshal(feat)
	if err != nil {
		log.Println(err)
	}

	return persistence.UpdateFeature(jsonData)
}

// ToggleFeature ...
func ToggleFeature(id int) bool {
	return persistence.ToggleFeature(id)
}

// ArchiveFeature ...
func ArchiveFeature(id int) bool {
	return persistence.ArchiveFeature(id)
}
