package model

import (
	"github.com/allegro/bigcache"
)

type App struct {
	InternalCache
	// Zookeeper
}

type InternalCache bigcache.BigCache

// type Zookeeper *zookeeperctl.Zookeeper
