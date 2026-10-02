package raft

import "sync"

type State int

const (
	Follower State = iota
	Candidate
	Leader
)

type LogEntry struct {
	Term  int
	Key   string
	Value string
}

type Node struct {
	mu sync.Mutex

	ID int

	CurrentTerm int
	VotedFor    int

	State State

	Log []LogEntry
}

func NewNode(id int) *Node {
	return &Node{
		ID:          id,
		CurrentTerm: 0,
		VotedFor:    -1,
		State:       Follower,
		Log:         make([]LogEntry, 0),
	}
}

func (n *Node) BecomeCandidate() {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.State = Candidate
	n.CurrentTerm++
	n.VotedFor = n.ID
}

func (n *Node) BecomeLeader() {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.State = Leader
}

func (n *Node) BecomeFollower(term int) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.State = Follower
	n.CurrentTerm = term
	n.VotedFor = -1
}

func (n *Node) AppendEntry(key, value string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.Log = append(n.Log, LogEntry{
		Term:  n.CurrentTerm,
		Key:   key,
		Value: value,
	})
}
