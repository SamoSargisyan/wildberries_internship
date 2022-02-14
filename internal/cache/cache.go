package cache

import (
	"github.com/patrickmn/go-cache"
	"l0/internal/domain"
	"time"
)

type LocalCache struct {
	orders *cache.Cache
}

func NewCache(orders []domain.OrderEntity) *LocalCache {

	c := cache.New(1*time.Hour, 1*time.Hour)

	for i := range orders {
		c.Set(orders[i].OrderUID, orders[i], cache.NoExpiration)
	}

	return &LocalCache{
		orders: c,
	}
}

func (cacheLocal *LocalCache) Set(key string, value interface{}, duration time.Duration) {
	cacheLocal.orders.Set(key, value, duration)
}

func (cacheLocal *LocalCache) Get(key string) interface{} {
	order, _ := cacheLocal.orders.Get(key)
	if order == nil {
		panic("Error while receiving data from cache")
	}

	return order
}
