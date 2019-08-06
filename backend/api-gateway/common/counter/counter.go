package counter

import (
	"fmt"

	"github.com/thinhlvv/tinyurl/backend/api-gateway/common/internalcache"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/common/zookeeperctl"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/config"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/internal/model"
)

// Counter struct holds interfaces module Counter needs.
type Counter struct {
	cache     internalcache.Engine
	zookeeper zookeeperctl.Zookeeper
	config    *config.Config
}

// New ...
func New(app *model.App) *Counter {
	return &Counter{
		cache:     app.InternalCache,
		zookeeper: app.Zookeeper,
		config:    app.Config,
	}
}

// MustInit ...
func (c *Counter) MustInit() error {
	// check cache first
	_, err := c.GetOrderNumber()
	if err == nil {
		return nil
	}

	// Check zookeeper with server name.
	value, err := c.zookeeper.Read(fmt.Sprintf("/counter/%s", c.config.ServerName))
	if err == nil {
		err = c.cache.Set(internalcache.ORDER_NUMBER_KEY, string(value))
		if err != nil {
			fmt.Println("[SetOrderNumber] Error when setting order number")
			return err
		}
	}

	// if not create new one

	// if cache not existed -> check zookeeper
	// if zookeeper hasnt registered yet -> register and update cache
	// if zookeeper registered alr -> get order number

	return nil
}

func (c *Counter) GetOrderNumber() (int, error) {
	cache, err := c.cache.Get(internalcache.ORDER_NUMBER_KEY)
	if err != nil {
		fmt.Println("[GetOrderNumber] Error empty order number")
		return 0, err
	}

	orderNumber, err := cache.Int()
	if err != nil {
		fmt.Println("[GetOrderNumber] Error can't convert string to int:", err)
		return 0, err
	}
	return orderNumber, nil
}
