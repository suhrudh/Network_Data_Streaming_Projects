Name : Suhrudh Reddy Sannapareddy
UFID : 64851063

This project is written in Golang.

Hash Function:
The function used is inspired from the Knuth's Multiplicative hashing, Multiplying the flowID with the special numnber '2654435761' 
to produce a evenly hashed result set over 2^32 range.

Project structure:
1. Hash/hash.go  -> Contains the core logic for all the different hashes, same file as project 1.
2. CustomRandom/customrandom.go -> contains the implementation of random number genration logic using unix time, same file as project 1.
3. PreProcessing/preprocessing.go -> Does input parsing from the file and assigns random numbers as flowids to each IP address, creates several objects useful for project3.go
4. CountMin -> Contains core logic for Count Min and Counter Sketch
5. project3.go -> Contains the main driver functions for all three Counters implementations and core logic & implemenation of active counter.(entry point for code - main).

Steps to run
1. Go inside the directory, where project3.go is present.
2. command => go run project3.go

Please check the comment just above the respective functions in main, to change inputs of the functions. Set to demo inputs by default.

Outputs:
All the Outputs are in "output" directory.
1. active_counter_output.txt => output for active counter.
2. count_min_output.txt => output for count min filter.
3. counter_sketch_output.txt => output for counter sketch.

Sample Outputs Acheived are present in "output_recorded" directory for reference, these are for the demo inputs.