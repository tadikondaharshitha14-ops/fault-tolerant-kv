package raft

func Replicate(leader *Node, followers []*Node) int {
	leader.mu.Lock()

	entries := make([]LogEntry, len(leader.Log))
	copy(entries, leader.Log)

	term := leader.CurrentTerm
	leaderID := leader.ID

	leader.mu.Unlock()

	successes := 1

	for _, follower := range followers {
		response := follower.AppendEntries(AppendEntriesRequest{
			Term:     term,
			LeaderID: leaderID,
			Entries:  entries,
		})

		if response.Success {
			successes++
		}
	}

	return successes
}
