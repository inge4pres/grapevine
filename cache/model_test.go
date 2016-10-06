package cache

import (
	"reflect"
	"testing"
)

var c Cacher

func TestCreate(t *testing.T) {
	c := NewCache(CACHE_MEMORY)
	m := NewMmap()
	if reflect.TypeOf(c) != reflect.TypeOf(m) {
		t.Errorf("Different types of objects: %v vs. %v", c, m)
	}

	u := NewCache(4)
	if u != nil {
		t.Errorf("Expected nil, found %v", u)
	}
}
