package raft

type AppendEntriesRequest struct {
	Term     int
	LeaderID int
	Entries  []LogEntry
}

type AppendEntriesResponse struct {
	Term    int
	Success bool
}

func (n *Node) AppendEntries(req AppendEntriesRequest) AppendEntriesResponse {
	n.mu.Lock()
	defer n.mu.Unlock()

	if req.Term < n.CurrentTerm {
		return AppendEntriesResponse{
			Term:    n.CurrentTerm,
			Success: false,
		}
	}

	if req.Term > n.CurrentTerm {
		n.CurrentTerm = req.Term
	}

	n.State = Follower
	n.VotedFor = -1

	n.Log = append(n.Log, req.Entries...)

	return AppendEntriesResponse{
		Term:    n.CurrentTerm,
		Success: true,
	}
}
