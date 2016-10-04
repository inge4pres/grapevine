package cache

func NewMmap() *mmap {
	return &mmap{
		desc: make(map[string]*mem, 0),
	}
}

func (m mmap) Get(key string) []byte {
	//FIXME use btree iterator
	for k, p := range m.desc {
		if k == key {
			return p.data
		}
	}
	return nil
}

func (m mmap) Set(key string, cont []byte, ttl int64) error {
	m.desc[key] = &mem{
		data: cont,
		ttl:  ttl,
	}
	//TODO ttl logic, config params
	return nil
}

func (m mmap) Flush() {
	m.desc = make(map[string]*mem, 0)
}

func (m mmap) Keys() []string {
	ret := make([]string, 0)
	for k, _ := range m.desc {
		ret = append(ret, k)
	}
	return ret
}
