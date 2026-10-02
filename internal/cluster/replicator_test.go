package cluster

import "testing"

func TestReplication(t *testing.T) {
	node1 := NewNode(1, ":50051")
	node2 := NewNode(2, ":50052")
	node3 := NewNode(3, ":50053")

	replicator := NewReplicator(
		node1,
		node2,
		node3,
	)

	replicator.Set("name", "harshitha")

	for _, node := range []*Node{node1, node2, node3} {
		value, ok := node.Get("name")

		if !ok {
			t.Fatalf("node %d did not receive replicated value", node.ID)
		}

		if value != "harshitha" {
			t.Fatalf(
				"node %d: expected harshitha, got %s",
				node.ID,
				value,
			)
		}
	}
}

func TestReplicationDelete(t *testing.T) {
	node1 := NewNode(1, ":50051")
	node2 := NewNode(2, ":50052")
	node3 := NewNode(3, ":50053")

	replicator := NewReplicator(
		node1,
		node2,
		node3,
	)

	replicator.Set("course", "MTECH CSE")
	replicator.Delete("course")

	for _, node := range []*Node{node1, node2, node3} {
		_, ok := node.Get("course")

		if ok {
			t.Fatalf("node %d still has deleted key", node.ID)
		}
	}
}
