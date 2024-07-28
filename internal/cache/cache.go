package cache

import (
	"github.com/erni27/imcache"
)

func NewTTLClient[K comparable, V any](limit int) *TTLClient[K, V] {
	cli := imcache.New[K, V](imcache.WithMaxEntriesLimitOption[K, V](limit, imcache.EvictionPolicyLFU))

	return &TTLClient[K, V]{
		cli: cli,
	}
}

type TTLClient[K comparable, V any] struct {
	cli *imcache.Cache[K, V]
}

func (c *TTLClient[K, V]) Get(key K) (V, bool) {
	return c.cli.Get(key)
}

func (c *TTLClient[K, V]) Set(key K, value V) {
	c.cli.Set(key, value, nil)
}
func (c *TTLClient[K, V]) Delete(key K) {
	c.cli.Remove(key)
}
func (c *TTLClient[K, V]) GetAll() map[K]V {
	return c.cli.GetAll()
}
