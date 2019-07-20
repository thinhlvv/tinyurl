package model

import (
	"github.com/thinhlvv/tinyurl/backend/common/internalcache"
	"github.com/thinhlvv/tinyurl/backend/common/zookeeperctl"
)

type App struct {
	InternalCache
	Zookeeper
}

type InternalCache internalcache.Engine
type Zookeeper zookeeperctl.Zookeeper
