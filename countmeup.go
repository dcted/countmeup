package main

import (
	"os"
	"fmt"
	"log"
)

type Votes struct {
	//to keep data to a minimum, our candidate numbers will be uint8 (0-255) assuming we have no more than 256 "candidates"
	votes []uint8
}

//The main data set read in from file
var voters = make(map[string]Votes)

//stores tally of votes for candidates as votes are counted
var candidateTotals = make(map[uint8]int)
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
}