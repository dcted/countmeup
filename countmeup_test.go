package main

import (
	"testing"
)

//Desired: Max three votes. should only have three votes for user
func TestUserMoreThanThreeVotes(t *testing.T) {
	insertVote("1","1")
	insertVote("1","2")
	insertVote("1","3")
	insertVote("1","4")
	insertVote("1","5")
	
	if(len(voters["1"].votes) != 3) {
		t.Fail()
	}
}

//Desired: three votes for user
func TestUserExactlyThreeVotes(t *testing.T) {
	insertVote("2","1")
	insertVote("2","2")
	insertVote("2","3")
	
	if(len(voters["1"].votes) != 3) {
		t.Fail()
	}
}

//Desired: user should have no votes
func TestNonExistingUser(t *testing.T) {
	insertVote("3","1")
	insertVote("3","2")
	
	if(len(voters["4"].votes) != 0) {
		t.Fail()
	}
}

func TestNonExistingCandidate(t *testing.T) {
	if(candidateTotals[6] != 0) {
		t.Fail()
	}
}

//Desired: 1 should have 50% of the votes, 6 total
// 2 should have 25% of the votes, 3 total
// 3 should have 35% of the votes, 3 total
func TestCandidateVoteTallies(t *testing.T) {
	insertVote("5","6")
	insertVote("5","6")
	insertVote("5","6")
	insertVote("5","6") //should void
	insertVote("5","6") //should void
	insertVote("6","6")
	insertVote("6","6")
	insertVote("6","6")
	insertVote("6","6") //should void
	insertVote("7","7")
	insertVote("7","7")
	insertVote("7","7")
	insertVote("8","8")
	insertVote("8","8")
	insertVote("8","8")
	insertVote("8","8") //should void
	
	if(candidateTotals[6] != 6 || candidateTotals[7] != 3 || candidateTotals[8] != 3) {
		t.Fail()
	}
}