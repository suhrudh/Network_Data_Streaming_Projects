Name : Suhrudh Reddy Sannapareddy
UFID : 64851063

This project is written in Golang.

Hash Function:
The function used is inspired from the Knuth's Multiplicative hashing, Multiplying the flowID with the special numnber '2654435761' 
to produce a evenly hashed result set over 2^32 range.

Project structure:
1. Hash/hash.go  -> Contains the core logic for all the different hashes.
2. CustomRandom/customrandom.go -> contains the implementation of random number genration logic using unix time.
3. Stream/stream.go -> Has the logic for stream generation(used for flowIDs generation).
4. project1.go -> Contains the main deiver functions for all three Hash table implementations.(entry point for code - main).

Steps to run
1. Go inside the directory, where project1.go is present.
2. command => go run project1.go

Please check the comment just above the respective functions in main, to change inputs of the functions. Set to demo inputs by default.

Outputs:
All the Outputs are in "output" directory.
1. cuckoo_output.txt => output for cuckoo hashing.
2. dleft_output.txt => output for dleft hashing.
3. multi_output.txt => output for multi hashing.

Sample Outputs Acheived are present in "output_recorded" directory for reference, these are for the demo inputs.