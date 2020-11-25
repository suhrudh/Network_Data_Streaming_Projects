Name : Suhrudh Reddy Sannapareddy
UFID : 64851063

This project is written in Golang.

Hash Function:
The function used is inspired from the Knuth's Multiplicative hashing, Multiplying the flowID with the special numnber '2654435761' 
to produce a evenly hashed result set over 2^32 range.

Project structure:
1. Hash/hash.go  -> Contains the core logic for all the different hashes, same file as project 1.
2. CustomRandom/customrandom.go -> contains the implementation of random number genration logic using unix time, same file as project 1.
3. PreProcessing/preprocessing.go -> Does input parsing from the file and assigns random numbers as flowids to each IP address, creates several objects useful for project4.go
4. set/set.go -> has logic for generating unique ids in the set.
5. project3.go -> Contains the logic for virtual bitmap simulation for given input traffic.(entry point for code - main).

Steps to run
1. Go inside the directory, where project4.go is present.
2. command => go run project4.go

Please check the comment just above the respective functions in main, to change inputs of the functions. Set to demo inputs by default.

Outputs:
All the Outputs are in "output" directory.
1. x_values.txt => extracted input from "project4input.txt" file.
2. y_values.txt => estimated output values.

=> Please find multi_flow_spread_size_comparision_graph.pdf which contains the plot.