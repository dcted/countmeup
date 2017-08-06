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

		//TODO first try and retrieve the voter, if theyre not there, we'll insert them with their vote
		//if they are there we will check their votes and if they have < 3 votes, we will add the vote, else skip
	}
	
	elapsed := time.Since(start)
	fmt.Println("Processing took ", elapsed)
}