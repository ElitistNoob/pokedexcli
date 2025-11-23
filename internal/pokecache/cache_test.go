package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCacheAdd(t *testing.T) {
	const interval = 5 * time.Minute

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "www.example.com",
			val: []byte("Some Data"),
		},
		{
			key: "www.2ndexample.com",
			val: []byte("Some Other Data"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected key not found")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("Value does not match expected value")
				return
			}
		})
	}
}

func TestReadLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)

	cache.Add("www.example.com", []byte("Some Data"))
	_, ok := cache.Get("www.example.com")
	if !ok {
		t.Errorf("Expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("www.example.com")
	if ok {
		t.Errorf("Expected not to find key")
		return
	}

}
