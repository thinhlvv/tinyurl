package internalcache

import (
	"log"
	"time"

	"github.com/allegro/bigcache"
)

type Engine struct {
	cache *bigcache.BigCache
}

func New() *Engine {
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))

	if err != nil {
		log.Fatal("Can't define internal cache engine:", err)
	}

	return &Engine{
		cache: cache,
	}
}
