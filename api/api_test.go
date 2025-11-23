package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestApi(t *testing.T) {
	expected := Locations{
		Count: 1,
		Results: []Result{
			{Name: "pallet-town", Url: "https://pokeapi.co/api/v2/location-area/1/"},
		},
	}

	mockData, _ := json.Marshal(expected)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mockData)
	}))
	defer server.Close()

	c := NewRequest(5*time.Second, time.Minute*5)

	locations, err := c.GetLocations(&server.URL)
	if err != nil {
		t.Errorf("Expected no errors %v", err)
		return
	}

	if locations.Count == 0 {
		t.Errorf("Expected count: %d, got %d", expected.Count, locations.Count)
	}

	if expected.Results[0].Name != locations.Results[0].Name {
		t.Errorf("Expected Name: %s, got: %s", expected.Results[0].Name, locations.Results[0].Name)
	}

}
