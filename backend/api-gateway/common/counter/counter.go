package counter

import (
	"fmt"

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

func (c *Counter) MustInit() error {
	// check cache first
	orderNumber, err := c.GetNumber()
	if err != nil {
		return err
	}
	if orderNumber > 0 {
		return nil
	}

	// if cache not existed -> check zookeeper
	// if zookeeper hasnt registered yet -> register and update cache
	// if zookeeper registered alr -> get order number

	return nil
}

func (c *Counter) GetNumber() (int, error) {
	cache, err := c.cache.Get(internalcache.ORDER_NUMBER_KEY)
	orderNumber, err := cache.Int()
	if err != nil {
		fmt.Println("[GetNumber] Error can't convert string to int:", err)
		return 0, err
	}
	return orderNumber, nil
}
