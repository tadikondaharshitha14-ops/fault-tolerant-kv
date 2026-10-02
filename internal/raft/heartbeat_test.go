package raft

import "testing"

func TestHeartbeat(t *testing.T) {
	leader := NewNode(1)
	follower := NewNode(2)

	leader.BecomeCandidate()
	leader.BecomeLeader()

	successes := SendHeartbeats(
		leader,
		[]*Node{follower},
	)

	if successes != 2 {
		t.Fatalf(
			"expected 2 alive nodes, got %d",
			successes,
		)
	}

	if follower.State != Follower {
		t.Fatal("follower should remain follower")
	}

	if follower.CurrentTerm != leader.CurrentTerm {
		t.Fatal("follower term should match leader term")
	}
}

func TestHeartbeatUpdatesHigherTerm(t *testing.T) {
	leader := NewNode(1)
	follower := NewNode(2)

	leader.BecomeCandidate()
	leader.BecomeLeader()

	leader.mu.Lock()
	leader.CurrentTerm = 5
	leader.mu.Unlock()

	response := follower.Heartbeat(HeartbeatRequest{
		Term:     5,
		LeaderID: 1,
	})

	if !response.Alive {
		t.Fatal("heartbeat should be accepted")
	}

	if follower.CurrentTerm != 5 {
		t.Fatalf(
			"expected term 5, got %d",
			follower.CurrentTerm,
		)
	}
}

func TestOldHeartbeatRejected(t *testing.T) {
	follower := NewNode(2)

	follower.mu.Lock()
	follower.CurrentTerm = 5
	follower.mu.Unlock()

	response := follower.Heartbeat(HeartbeatRequest{
		Term:     3,
		LeaderID: 1,
	})

	if response.Alive {
		t.Fatal("old heartbeat should be rejected")
	}
}
