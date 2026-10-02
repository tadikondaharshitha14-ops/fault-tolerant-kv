package cluster

import (
	"sync"
)

type Replicator struct {
	mu    sync.RWMutex
	nodes []*Node
}

func NewReplicator(nodes ...*Node) *Replicator {
	return &Replicator{
		nodes: nodes,
	}
}

func (r *Replicator) Set(key, value string) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, node := range r.nodes {
		node.Set(key, value)
	}
}

func (r *Replicator) Get(key string) (string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.nodes) == 0 {
		return "", false
	}

	return r.nodes[0].Get(key)
}

func (r *Replicator) Delete(key string) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, node := range r.nodes {
		node.Delete(key)
	}
}
