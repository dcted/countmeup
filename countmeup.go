package main

// To start, thinking about data structures to hold data in the way described, it made the most sense to use a map, as it is a fast add 
// and fast look up which is most crutial for this system

// It was obvious that file IO was going to be the main time consuming part of each run. Part of the assumption then is how the votes will come to the system.
// Are they all collated? Are the coming in real time? - If they are all collated already into a file, which is the assumption I have gone with, then the IO
// of the file is the biggest issue. The data could be held in memory for subsequent runs, focusing on the diff between the new and old data.
// Part of the idea of going with go was knowing how well the language can parallelise tasks. The idea being to branch the logic off straight after reading 
// from file so that the throttle is almost solely on the IO scan.

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"strconv"
	"strings"
	"time"
	//"sync"
)

type Votes struct {
	//to keep data to a minimum, our candidate numbers will be uint8 (0-255) assuming we have no more than 256 "candidates"
	votes []uint8
}

//The main data set read in from file
var voters = make(map[string]Votes)
//var votersLock = &sync.Mutex{}

//stores tally of votes for candidates as votes are counted
var candidateTotals = make(map[uint8]int)
//var candidateTotalsLock = &sync.Mutex{}
var totalVotes = 0

func main() {
	//TODO: perhaps add argument to create data set for testing purposes
	if len(os.Args) != 2 {
		fmt.Println("Run this by including the data file path as the first agrument")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	parseVotes(file)
}

//Working assumption that the file is local and this program will be run with the entire data set included each time
//this isnt likely in practice but as a start we will go with this idea
func parseVotes(file *os.File) {
	scan := bufio.NewScanner(file)

	start := time.Now()
	//Working on assumption that data is in some Id:vote key value fashion (even inherently as in, phonenumber + vote)
	//and that it is tab separated and 1 vote per line
	for scan.Scan() {
		row := strings.Split(scan.Text(), "\t")

		//just continuing may not be the best approach, but for now
		if len(row) != 2 {
			continue
		}

		//first try and retrieve the voter, if theyre not there, we'll insert them with their vote
		//if they are there we will check their votes and if they have < 3 votes, we will add the vote, else skip
		//go insertVote(row[0], row[1])
		insertVote(row[0], row[1])
	}
	elapsed := time.Since(start)
	fmt.Println("Processing took ", elapsed)

	tallyVotes()
}

func insertVote(voter string, vote string) {
	voteForCandidate, err := strconv.ParseInt(vote, 0, 8)

	if err != nil {
		return
	}

	//If we have record of the voter, attempt add, otherwise just add
	//votersLock.Lock()
	votes, exists := voters[voter] 
	if len(votes.votes) < 3 {
		if exists {
			votes.votes = append(votes.votes, uint8(voteForCandidate))
		} 
		
		voters[voter] = votes
	}
	//votersLock.Unlock()

	//candidateTotalsLock.Lock()
	candidateTotals[uint8(voteForCandidate)] = candidateTotals[uint8(voteForCandidate)] + 1
	totalVotes++
	//candidateTotalsLock.Unlock()
}

func tallyVotes() {
	fmt.Println("candidate\ttotal")

	for key,val := range candidateTotals {
		perc := float64(val) / float64(totalVotes)
		fmt.Println(fmt.Sprintf("candidate %d %c %d ( %f )", key, '\t', val, perc))
	}
}