package cache

import (
	"os"
)

const (
	DEFAULT_TTL = 3600000
)

type Cacher interface {
	Get(string) []byte
	Set(string, interface{}, int64) error
	Flush() error
	Keys() []string
}

type mem struct {
	ttl  int64
	data []byte
}

type mmap struct {
	//TODO dd namespace to maps!
	desc map[string]*mem
}

type file struct {
	ttl  int64
	data []byte
}

type fmap struct {
	desc map[*os.FileInfo]*file
}

//TODO
//func btreeIterator() token.Token {
//	return nil
//}
