package model

import (
	"github.com/allegro/bigcache"
	"github.com/thinhlvv/tinyurl/backend/common/zookeeperctl"
)

type App struct {
	InternalCache
	Zookeeper
}

type InternalCache bigcache.BigCache

type Zookeeper zookeeperctl.Zookeeper
