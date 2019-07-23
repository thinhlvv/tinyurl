package internalcache

import (
	"log"
	"time"

	"github.com/allegro/bigcache"
)

type Engine struct {
	cache *bigcache.BigCache
}

func New() Engine {
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Second))

	if err != nil {
		log.Fatal("Can't define internal cache engine:", err)
	}

	return Engine{
		cache: cache,
	}
}

func (e *Engine) Set(key, value string) error {
	return e.cache.Set(key, []byte(value))
}

func (e *Engine) Get(key string) ([]byte, error) {
	return e.cache.Get(key)
}
