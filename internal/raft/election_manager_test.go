package raft

import "testing"

func TestThreeNodeElection(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)

	leader := StartElection(
		node1,
		[]*Node{node2, node3},
	)

	if !leader {
		t.Fatal("node 1 should win election")
	}

	if node1.State != Leader {
		t.Fatal("node 1 should become leader")
	}

	if node1.CurrentTerm != 1 {
		t.Fatalf(
			"expected term 1, got %d",
			node1.CurrentTerm,
		)
	}
}
