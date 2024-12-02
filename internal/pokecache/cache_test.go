package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cached == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cache.Add("Key", []byte("Value"))
	val, ok := cache.Get("Key")
	if !ok {
		t.Error("key not found")
	}
	if string(val) != "Value" {
		t.Error("value not found")
	}
}

func TestClearCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 10)
	cache.Add("Key", []byte("Value"))
	time.Sleep(time.Millisecond * 100)
	_, ok := cache.Get("Key")
	if ok {
		t.Error("key still found in cache")
	}
}

func TestClearCacheFail(t *testing.T) {
	cache := NewCache(time.Millisecond * 10)
	cache.Add("Key", []byte("Value"))
	time.Sleep(time.Millisecond)
	_, ok := cache.Get("Key")
	if !ok {
		t.Error("key not found in cache")
	}
}
