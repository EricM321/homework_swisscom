package main

import (
	"encoding/json"
	"fmt"
	"log"
	"persistence"
	"reflect"
	"testing"
	"time"
)

func init() {
	persistence.CreateDatabase()
	fmt.Println("\nTesting feature")
}

func TestGetFeatures(t *testing.T) {
	expected := []byte(`{"features":[{"featureId":1,"displayName":"firstFeature","technicalName":"testingFeature","expiresOn":"2022-12-31T12:00:00Z","description":"Checking if insert works","inverted":false,"active":true}]}`)
	actual := GetFeatures()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("GetFeatures was incorrect:\nactual:   %s\nexpected: %s", actual, expected)
	}
}

func TestGetFeature(t *testing.T) {
	expected := []byte(`{"featureId":1,"displayName":"firstFeature","technicalName":"testingFeature","expiresOn":"2022-12-31T12:00:00Z","description":"Checking if insert works","inverted":false,"active":true}`)
	actual := GetFeature(1)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("GetFeature was incorrect:\nactual:   %s\nexpected: %s", actual, expected)
	}
}

func TestCreateNewFeature(t *testing.T) {
	values := Feature{
		DisplayName:   "Test",
		TechnicalName: "Test_Create_Feature",
		ExpiresOn:     time.Time{}, // "empty" time value
		Description:   "",
		Inverted:      false,
		CustomerIds:   []int{1, 2},
	}

	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))

	expected := true //[]byte(`{"featureId":2,"displayName":"Test","technicalName":"Test_Create_Feature","expiresOn":"0001-01-01T00:00:00Z","description":"","inverted":false,"active":true}`)
	actual := CreateNewFeature(jsonData)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("CreateNewFeature was incorrect:\nactual:   %t\nexpected: %t", actual, expected)
	}
}

func TestUpdateFreature(t *testing.T) {
	values := Feature{
		ID:            2,
		DisplayName:   "Test",
		TechnicalName: "Test_Create_Feature",
		ExpiresOn:     time.Time{}, // "empty" time value
		Description:   "",
		Inverted:      false,
		CustomerIds:   []int{1},
	}

	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))

	expected := true
	actual := UpdateFreature(jsonData)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("UpdateFreature was incorrect:\nactual:   %t\nexpected: %t", actual, expected)
	}
}

func TestToggleFeature(t *testing.T) {
	expected := true
	actual := ToggleFeature(1)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ToggleFeature was incorrect:\nactual:   %t\nexpected: %t", actual, expected)
	}
}

func TestArchiveFeature(t *testing.T) {
	expected := true
	actual := ArchiveFeature(2)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ArchiveFeature was incorrect:\nactual:   %t\nexpected: %t", actual, expected)
	}
}
