package cache

import (
	"os"
)

const (
	CACHE_MEMORY = iota
	CACHE_FILE

	DEFAULT_TTL = 3600000
)

type Cacher interface {
	Get(string) []byte
	Set(string, []byte, int64) error
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
	desc map[os.FileInfo]*file
}

func NewCache(ctype int) Cacher {
	switch ctype {
	case CACHE_MEMORY:
		return NewMmap()
	case CACHE_FILE:
		return NewFmap()
	default:
		return nil
	}
	return nil
}

//TODO
//func btreeIterator() token.Token {
//	return nil
//}
