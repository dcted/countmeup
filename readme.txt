Count Me Up

Count My Up is written in Golang and therefore requires Go to be installed.

Can be run as usual from in shell or command line, with an added argument as 
the path to the file containing the votes. The votes txt file should be
in the format <VOTER_ID>\t<CANDIDATE_ID>\n for each line.

A file is included as votecreator.go in the folder votecreator to easily
create a txt file of votes for Count Me Up.

A test file is included as votecreator_test.go which can be run as a normal go test.