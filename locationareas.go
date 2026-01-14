package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/B00m3r0302/PokeDex/internal/pokecache"
)

const pageSize = 20

type MapNavigator struct {
	offset int
	total  int
	cache  *pokecache.Cache
}

type LocationAreaResponse struct {
	Count   int       `json:"count"`
	Results []NameURL `json:"results"`
}

type NameURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func NewMapNavigator() *MapNavigator {
	return &MapNavigator{
		cache: pokecache.NewCache(10 * time.Second),
	}
}

func (m *MapNavigator) Reset() {
	m.offset = 0
	m.total = 0
}

func (m *MapNavigator) LocationAreaMoveForward() error {
	page, err := GetLocationArea(m.offset)
	if err != nil {
		return err
	}

	if m.offset == 0 {
		m.total = page.Count
	}

	for _, loc := range page.Results {
		fmt.Println(loc.Name)
		m.cache.Add(loc.Name, []byte(loc.Name))
	}

	m.offset += pageSize

	if m.offset >= m.total {
		fmt.Println("Reached the end of the location areas.")
		m.Reset()
	}

	return nil
}

func (m *MapNavigator) LocationAreaMoveBackward() error {
	if m.offset < pageSize {
		fmt.Println("Already at the beginning of the location areas.")
		return nil
	}

	m.offset -= pageSize

	page, err := GetLocationArea(m.offset)
	if err != nil {
		return err
	}

	if m.total == 0 {
		m.total = page.Count
	}

	for _, loc := range page.Results {
		fmt.Println(loc.Name)
		m.cache.Add(loc.Name, []byte(loc.Name))
	}

	return nil
}

func GetLocationArea(offset int) (LocationAreaResponse, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?offset=%d&limit=20", offset)

	resp, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	var apiResp LocationAreaResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return LocationAreaResponse{}, err
	}

	return apiResp, nil
}
