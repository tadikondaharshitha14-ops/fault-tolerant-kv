package cluster

import (
	"fmt"
	"sync"

	"fault-tolerant-kv/internal/kv"
)

type Node struct {
	mu sync.RWMutex

	ID   int
	Addr string

	Store *kv.Store
}

func NewNode(id int, addr string) *Node {
	return &Node{
		ID:    id,
		Addr:  addr,
		Store: kv.NewStore(),
	}
}

func (n *Node) Set(key, value string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.Store.Set(key, value)
}

func (n *Node) Get(key string) (string, bool) {
	n.mu.RLock()
	defer n.mu.RUnlock()

	return n.Store.Get(key)
}

func (n *Node) Delete(key string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.Store.Delete(key)
}

func (n *Node) String() string {
	return fmt.Sprintf(
		"Node{id=%d, addr=%s}",
		n.ID,
		n.Addr,
	)
}
