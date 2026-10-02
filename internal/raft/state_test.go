package raft

import (
	"testing"
	"time"
)

func TestNodeStatus(t *testing.T) {
	status := NewNodeStatus()

	if !status.Alive {
		t.Fatal("new node should be alive")
	}

	if status.TimedOut() {
		t.Fatal("new node should not be timed out")
	}
}

func TestNodeStatusFailure(t *testing.T) {
	status := NewNodeStatus()

	status.Fail()

	if status.Alive {
		t.Fatal("failed node should not be alive")
	}

	if !status.TimedOut() {
		t.Fatal("failed node should be considered timed out")
	}
}

func TestNodeStatusRecovery(t *testing.T) {
	status := NewNodeStatus()

	status.Fail()
	status.Recover()

	if !status.Alive {
		t.Fatal("recovered node should be alive")
	}
}

func TestElectionTimeout(t *testing.T) {
	status := NewNodeStatus()

	status.LastHeartbeat = time.Now().Add(-500 * time.Millisecond)

	if !status.TimedOut() {
		t.Fatal("node should have timed out")
	}
}
