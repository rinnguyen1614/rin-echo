package groupcache

import "sync"

type GroupCache struct {
	mu sync.Mutex
}
