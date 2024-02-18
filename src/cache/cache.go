package cache

import (
	"github.com/patrickmn/go-cache"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/model"
	"log"
)

type AllCache struct {
	clients *cache.Cache
}

var c *AllCache

const (
	defaultExpiration = cache.NoExpiration
	cleanupInterval   = cache.NoExpiration
)

func init() {
	c = newCache()
	c.initializeClients()
}

func newCache() *AllCache {
	ClientCache := cache.New(defaultExpiration, cleanupInterval)
	return &AllCache{
		clients: ClientCache,
	}
}

func GetCache() *AllCache {
	return c
}

func (c *AllCache) ReadClient(id string) (item *model.User) {
	if client, found := c.clients.Get(id); found {
		if res, ok := client.(*model.User); ok {
			return res
		}
		log.Fatal("Error decoding cache value: Type is not User")
	}
	return nil
}

func (c *AllCache) UpdateClient(id string, newValue *model.User) {
	c.clients.Set(id, newValue, cache.NoExpiration)
}

func (c *AllCache) initializeClients() {
	client1 := &model.User{
		Balance: 0,
		Limit:   100000,
	}
	client2 := &model.User{
		Balance: 0,
		Limit:   80000,
	}
	client3 := &model.User{
		Balance: 0,
		Limit:   1000000,
	}
	client4 := &model.User{
		Balance: 0,
		Limit:   10000000,
	}
	client5 := &model.User{
		Balance: 0,
		Limit:   500000,
	}

	c.clients.Set("1", client1, cache.NoExpiration)
	c.clients.Set("2", client2, cache.NoExpiration)
	c.clients.Set("3", client3, cache.NoExpiration)
	c.clients.Set("4", client4, cache.NoExpiration)
	c.clients.Set("5", client5, cache.NoExpiration)
}
