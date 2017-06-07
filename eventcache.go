package monitaur

import (
	"sync"
)

type EventCache struct {
	*sync.RWMutex
	lookup map[string]History
}

func NewEventCache() *EventCache {
	return &EventCache{
		RWMutex: &sync.RWMutex{}, lookup: make(map[string]History),
	}
}

func (cache *EventCache) Save(event Event) {
	cache.Lock()
	defer cache.Unlock()
	key := keyOf(event.Client, event.Check)
	if _, ok := cache.lookup[key]; ok {
		cache.lookup[key] = append(cache.lookup[key], event)
	} else {
		cache.lookup[key] = make(History, 0, 1)
		cache.lookup[key] = append(cache.lookup[key], event)
	}
}

func (cache *EventCache) Get(client *Client, check Check) (History, bool) {
	cache.RLock()
	defer cache.RUnlock()
	hist, ok := cache.lookup[keyOf(client, check)]
	return hist, ok
}

func keyOf(client *Client, check Check) string {
	return client.Name + check.GetName()
}
