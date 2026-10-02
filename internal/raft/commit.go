package raft

func Majority(total int) int {
	return total/2 + 1
}

func HasQuorum(totalNodes, successfulNodes int) bool {
	return successfulNodes >= Majority(totalNodes)
}

func (n *Node) CommitEntry(index int) bool {
	n.mu.Lock()
	defer n.mu.Unlock()

	if index < 0 || index >= len(n.Log) {
		return false
	}

	return true
}
