package raft

type RequestVoteRequest struct {
	Term         int
	CandidateID  int
	LastLogIndex int
	LastLogTerm  int
}

type RequestVoteResponse struct {
	Term        int
	VoteGranted bool
}

func (n *Node) RequestVote(req RequestVoteRequest) RequestVoteResponse {
	n.mu.Lock()
	defer n.mu.Unlock()

	if req.Term < n.CurrentTerm {
		return RequestVoteResponse{
			Term:        n.CurrentTerm,
		}
	}

	if req.Term > n.CurrentTerm {
		n.CurrentTerm = req.Term
		n.State = Follower
		n.VotedFor = -1
	}

	if n.VotedFor == -1 || n.VotedFor == req.CandidateID {
		n.VotedFor = req.CandidateID

		return RequestVoteResponse{
			Term:        n.CurrentTerm,
			VoteGranted: true,
		}
	}

	return RequestVoteResponse{
		Term:        n.CurrentTerm,
		VoteGranted: false,
	}
}
