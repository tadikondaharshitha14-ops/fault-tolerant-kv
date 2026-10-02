package raft

func StartElection(candidate *Node, peers []*Node) bool {
	candidate.BecomeCandidate()

	votes := 1

	for _, peer := range peers {
		response := peer.RequestVote(RequestVoteRequest{
			Term:        candidate.CurrentTerm,
			CandidateID: candidate.ID,
		})

		if response.VoteGranted {
			votes++
		}
	}

	majority := (len(peers)+1)/2 + 1

	if votes >= majority {
		candidate.BecomeLeader()
		return true
	}

	return false
}
