package pokecache

import (
	"testing"
	"time"
)

func TestPokecache(t *testing.T) {
	cache, _ := NewCache(1 * time.Second)
	cache.Add("foo", []byte("bar"))

	entry, _ := cache.Get("foo")
	if string(entry.val) != "bar" {
		t.Errorf("expected 'bar', got %s", string(entry.val))
	}

	time.Sleep(2 * time.Second)

	entry, err := cache.Get("foo")
	if err == nil {
		t.Errorf("expected nil, got %s", string(entry.val))
	}
}
