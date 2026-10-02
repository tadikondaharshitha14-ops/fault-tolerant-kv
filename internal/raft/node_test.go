package raft

import "testing"

func TestNewNode(t *testing.T) {
	node := NewNode(1)

	if node.ID != 1 {
		t.Fatalf("expected ID 1, got %d", node.ID)
	}

	if node.State != Follower {
		t.Fatal("new node should start as follower")
	}

	if node.CurrentTerm != 0 {
		t.Fatal("initial term should be 0")
	}
}

func TestBecomeCandidate(t *testing.T) {
	node := NewNode(1)

	node.BecomeCandidate()

	if node.State != Candidate {
		t.Fatal("node should become candidate")
	}

	if node.CurrentTerm != 1 {
		t.Fatalf("expected term 1, got %d", node.CurrentTerm)
	}

	if node.VotedFor != 1 {
		t.Fatalf("expected node to vote for itself")
	}
}

func TestBecomeLeader(t *testing.T) {
	node := NewNode(1)

	node.BecomeCandidate()
	node.BecomeLeader()

	if node.State != Leader {
		t.Fatal("node should become leader")
	}
}

func TestAppendEntry(t *testing.T) {
	node := NewNode(1)

	node.BecomeCandidate()
	node.BecomeLeader()

	node.AppendEntry("name", "harshitha")

	if len(node.Log) != 1 {
		t.Fatalf("expected 1 log entry, got %d", len(node.Log))
	}

	if node.Log[0].Key != "name" {
		t.Fatalf("expected key name, got %s", node.Log[0].Key)
	}

	if node.Log[0].Value != "harshitha" {
		t.Fatalf(
			"expected value harshitha, got %s",
			node.Log[0].Value,
		)
	}
}
