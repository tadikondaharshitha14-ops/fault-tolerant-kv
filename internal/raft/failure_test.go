package raft

import "testing"

func TestLeaderFailureAndReElection(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)

	// Node 1 becomes leader.
	node1.BecomeCandidate()
	node1.BecomeLeader()

	if node1.State != Leader {
		t.Fatal("node 1 should be leader")
	}

	// Node 1 fails.
	node1.Fail()

	// Node 2 starts a new election.
	won := StartElection(
		node2,
		[]*Node{node3},
	)

	if !won {
		t.Fatal("node 2 should win the new election")
	}

	if node2.State != Leader {
		t.Fatal("node 2 should become leader")
	}

	if node2.CurrentTerm != 1 {
		t.Fatalf(
			"expected new term 1, got %d",
			node2.CurrentTerm,
		)
	}
}
