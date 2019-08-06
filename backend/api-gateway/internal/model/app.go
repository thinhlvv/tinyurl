package model

import (
	"github.com/thinhlvv/tinyurl/backend/api-gateway/common/internalcache"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/common/zookeeperctl"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/config"
)

type App struct {
	InternalCache internalcache.Engine
	Zookeeper     zookeeperctl.Zookeeper
	Config        *config.Config
}
