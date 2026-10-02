package raft

func (n *Node) Fail() {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.State = Follower
	n.VotedFor = -1
}

func (n *Node) IsAlive() bool {
	n.mu.Lock()
	defer n.mu.Unlock()

	return n.CurrentTerm >= 0
}
