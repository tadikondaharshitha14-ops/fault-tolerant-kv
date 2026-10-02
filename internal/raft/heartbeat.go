package raft

type HeartbeatRequest struct {
	Term     int
	LeaderID int
}

type HeartbeatResponse struct {
	Term    int
	Alive   bool
}

func (n *Node) Heartbeat(req HeartbeatRequest) HeartbeatResponse {
	n.mu.Lock()
	defer n.mu.Unlock()

	if req.Term < n.CurrentTerm {
		return HeartbeatResponse{
			Term:  n.CurrentTerm,
			Alive: false,
		}
	}

	if req.Term > n.CurrentTerm {
		n.CurrentTerm = req.Term
	}

	n.State = Follower
	n.VotedFor = req.LeaderID

	return HeartbeatResponse{
		Term:  n.CurrentTerm,
		Alive: true,
	}
}
