package cache

import (
	"testing"
)

var test = []byte("A reasonable test content, a reasonable test result!")
var mm = NewMmap()

func TestSet(t *testing.T) {
	if err := mm.Set("tk-1", test, DEFAULT_TTL); err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	mem := mm.Get("tk-1")
	if len(mem) != len(test) {
		t.Errorf("Bad data in memory: %v vs. %v", mem, test)
	}
}
