package cache

import (
	"time"
)

func NewMmap() *mmap {
	return &mmap{
		desc: make(map[string]*mem, 0),
	}
}

// Get the corresponding data for the provided key
// The key gets expired by a separate process running in background
// Get returns nil either for a non-present key an for a expired key
func (m mmap) Get(key string) []byte {
	for k, p := range m.desc {
		if k == key {
			return p.data
		}
	}
	return nil
}

func (m mmap) Set(key string, cont []byte, ttl int64) error {
	live := int64(DEFAULT_TTL)
	if ttl > 0 {
		live = ttl
	}
	m.desc[key] = &mem{
		start: time.Now(),
		data:  cont,
		ttl:   live,
	}
	return nil
}

func (m mmap) Flush() error {
	m.desc = make(map[string]*mem, 0)
	return nil
}

func (m mmap) Keys() []string {
	ret := make([]string, 0)
	for k, _ := range m.desc {
		ret = append(ret, k)
	}
	return ret
}

func (m mmap) expire() {
	expCh := make(chan (*mem), 0)
	done := make(chan (bool))

	go func() {
		for _, o := range m.desc {
			if time.Since(o.start) > time.Duration(o.ttl) {
				expCh <- o
			}
		}
		done <- true
	}()
	for {
		select {
		case exp := <-expCh:
			exp.data = nil
		case <-done:
			m.expire()
		}
	}
	return
}
