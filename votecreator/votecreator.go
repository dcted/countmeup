package cmucreator

import (
    "math/rand"
    "time"
    "bufio"
    "os"
    "fmt"
)

var numberOfVotes = 10000000
var numberOfCandidates = 6
var numberOfVoters = 9000000

var currentVoter = 1;
var currentVote = 0;

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {
	// open output file
    fo, err := os.Create("votes.txt")
    if err != nil {
        panic(err)
    }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()

    // make a write buffer
    w := bufio.NewWriter(fo)

	for currentVote <= numberOfVotes {
		candidate := random(1, numberOfCandidates)

		if currentVoter >= numberOfVoters {
			currentVoter = 1
		} else {
			currentVoter++
		}

		//fmt.Println(fmt.Sprintf("%d%c%d", currentVoter, '\t', candidate))
		if _, err := w.WriteString(fmt.Sprintf("%d%c%d%c", currentVoter, '\t', candidate, '\n')); err != nil {
            panic(err)
        }

        w.Flush()

        currentVote++
	}
}