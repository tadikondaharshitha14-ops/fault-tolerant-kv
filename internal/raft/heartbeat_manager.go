package raft

func SendHeartbeats(leader *Node, followers []*Node) int {
	leader.mu.Lock()

	term := leader.CurrentTerm
	leaderID := leader.ID

	leader.mu.Unlock()

	successes := 1

	for _, follower := range followers {
		response := follower.Heartbeat(HeartbeatRequest{
			Term:     term,
			LeaderID: leaderID,
		})

		if response.Alive {
			successes++
		}
	}

	return successes
}
