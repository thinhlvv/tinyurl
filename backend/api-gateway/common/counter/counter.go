package counter

import (
	"github.com/thinhlvv/tinyurl/backend/api-gateway/common/internalcache"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/common/zookeeperctl"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/internal/model"
)

// This counter is next hash seed keeper
type Counter struct {
	cache     internalcache.Engine
	zookeeper zookeeperctl.Zookeeper
}

func New(app *model.App) *Counter {
	return &Counter{
		cache:     app.InternalCache,
		zookeeper: app.Zookeeper,
	}
}

func (c *Counter) GetNumber() uint64 {
	return 999
}
