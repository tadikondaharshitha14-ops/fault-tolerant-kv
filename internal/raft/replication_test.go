package raft

import "testing"

func TestLogReplication(t *testing.T) {
	leader := NewNode(1)
	follower2 := NewNode(2)
	follower3 := NewNode(3)

	leader.BecomeCandidate()
	leader.BecomeLeader()

	leader.AppendEntry("name", "harshitha")

	successes := Replicate(
		leader,
		[]*Node{
			follower2,
			follower3,
		},
	)

	if successes != 3 {
		t.Fatalf(
			"expected 3 successful nodes, got %d",
			successes,
		)
	}

	if len(follower2.Log) != 1 {
		t.Fatalf(
			"follower 2 expected 1 log entry, got %d",
			len(follower2.Log),
		)
	}

	if len(follower3.Log) != 1 {
		t.Fatalf(
			"follower 3 expected 1 log entry, got %d",
			len(follower3.Log),
		)
	}

	if follower2.Log[0].Value != "harshitha" {
		t.Fatalf("follower 2 has incorrect value")
	}

	if follower3.Log[0].Value != "harshitha" {
		t.Fatalf("follower 3 has incorrect value")
	}
}
