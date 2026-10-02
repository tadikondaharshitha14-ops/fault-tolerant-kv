package raft

import "time"

type NodeStatus struct {
	Alive             bool
	LastHeartbeat     time.Time
	ElectionTimeout   time.Duration
}

func NewNodeStatus() NodeStatus {
	return NodeStatus{
		Alive:           true,
		LastHeartbeat:   time.Now(),
		ElectionTimeout: 200 * time.Millisecond,
	}
}

func (s *NodeStatus) Touch() {
	s.LastHeartbeat = time.Now()
}

func (s *NodeStatus) Fail() {
	s.Alive = false
}

func (s *NodeStatus) Recover() {
	s.Alive = true
	s.Touch()
}

func (s *NodeStatus) TimedOut() bool {
	if !s.Alive {
		return true
	}

	return time.Since(s.LastHeartbeat) > s.ElectionTimeout
}
