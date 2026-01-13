package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(5 * time.Second)

	key := "testkey"
	value := []byte("testvalue")

	cache.Add(key, value)

	retrieved, ok := cache.Get(key)
	if !ok {
		t.Errorf("Expected to find key '%s' in cache, but it was not found", key)
	}

	if string(retrieved) != string(value) {
		t.Errorf("Expected value '%s', but got '%s'", string(value), string(retrieved))
	}
}

func TestGetNonExistent(t *testing.T) {
	cache := NewCache(5 * time.Second)

	_, ok := cache.Get("nonexistent")
	if ok {
		t.Errorf("Expected key 'nonexistent' to not be found, but it was")
	}
}

func TestReapLoop(t *testing.T) {
	interval := 100 * time.Millisecond
	cache := NewCache(interval)

	key := "tempkey"
	value := []byte("tempvalue")

	cache.Add(key, value)

	// Verify it exists immediately
	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("Expected key '%s' to exist immediately after adding", key)
	}

	// Wait for longer than the interval to allow reapLoop to clean it
	time.Sleep(interval + 50*time.Millisecond)

	// Now it should be gone
	_, ok = cache.Get(key)
	if ok {
		t.Errorf("Expected key '%s' to be reaped after interval, but it still exists", key)
	}
}
