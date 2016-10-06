package cache

import (
	"testing"
)

// table-driven tests ftw
var test = []byte("A reasonable test content, a reasonable test result!")
var objs = []struct {
	key  string
	cont []byte
	ttl  int64
}{
	{"k1", []byte("A reasonable test content, a reasonable test result!"), 100},
	{"k2", []byte("ABCDEFG"), 500},
	{"k3", []byte("lrnfq84hr8qweuidn034j3985ht24fg6h9230d1j278rhfowedw9wefon"), 1500},
}

var mm = NewMmap()

func TestSet(t *testing.T) {
	for _, o := range objs {
		if err := mm.Set(o.key, o.cont, o.ttl); err != nil {
			t.Errorf("Set() failed for %v %v %v", o.key, o.cont, o.ttl)
		}
	}
}

func TestGet(t *testing.T) {
	if mem := mm.Get("k1"); len(mem) != len(test) {
		t.Errorf("Bad data in memory: %v vs. %v", mem, test)
	}
	if mem := mm.Get("k4"); mem != nil {
		t.Error("Data from another key!")
	}
}

func TestKeys(t *testing.T) {
	if len(mm.Keys()) != len(objs) {
		t.Errorf("Lenght of keys differ from table tests, want %d have %d", len(objs), len(mm.Keys()))
	}
}
