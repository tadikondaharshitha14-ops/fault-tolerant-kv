package raft

import "testing"

func TestMajority(t *testing.T) {
	if Majority(3) != 2 {
		t.Fatalf("expected majority 2")
	}

	if Majority(5) != 3 {
		t.Fatalf("expected majority 3")
	}
}

func TestQuorum(t *testing.T) {
	if !HasQuorum(3, 2) {
		t.Fatal("2 of 3 should have quorum")
	}

	if HasQuorum(3, 1) {
		t.Fatal("1 of 3 should not have quorum")
	}

	if HasQuorum(5, 2) {
		t.Fatal("2 of 5 should not have quorum")
	}

	if !HasQuorum(5, 3) {
		t.Fatal("3 of 5 should have quorum")
	}
}

func TestCommitEntry(t *testing.T) {
	node := NewNode(1)

	node.BecomeCandidate()
	node.BecomeLeader()

	node.AppendEntry("name", "harshitha")

	if !node.CommitEntry(0) {
		t.Fatal("expected log entry to be committable")
	}

	if node.CommitEntry(10) {
		t.Fatal("invalid log index should not commit")
	}
}
