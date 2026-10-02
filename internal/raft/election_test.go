package raft

import "testing"

func TestRequestVote(t *testing.T) {
	node := NewNode(2)

	response := node.RequestVote(RequestVoteRequest{
		Term:        1,
		CandidateID: 1,
	})

	if !response.VoteGranted {
		t.Fatal("expected vote to be granted")
	}

	if response.Term != 1 {
		t.Fatalf("expected term 1, got %d", response.Term)
	}

	if node.VotedFor != 1 {
		t.Fatalf("expected vote for node 1, got %d", node.VotedFor)
	}
}

func TestRejectSecondCandidate(t *testing.T) {
	node := NewNode(2)

	first := node.RequestVote(RequestVoteRequest{
		Term:        1,
		CandidateID: 1,
	})

	if !first.VoteGranted {
		t.Fatal("first vote should be granted")
	}

	second := node.RequestVote(RequestVoteRequest{
		Term:        1,
		CandidateID: 3,
	})

	if second.VoteGranted {
		t.Fatal("second candidate should not receive another vote")
	}
}

func TestHigherTermUpdatesNode(t *testing.T) {
	node := NewNode(2)

	response := node.RequestVote(RequestVoteRequest{
		Term:        5,
		CandidateID: 1,
	})

	if !response.VoteGranted {
		t.Fatal("vote should be granted for higher term")
	}

	if node.CurrentTerm != 5 {
		t.Fatalf("expected term 5, got %d", node.CurrentTerm)
	}

	if node.State != Follower {
		t.Fatal("node should remain follower")
	}
}
