Name : Suhrudh Reddy Sannapareddy

This project is written in Golang.

Hash Function:
The function used is inspired from the Knuth's Multiplicative hashing, Multiplying the flowID with the special numnber '2654435761' 
to produce a evenly hashed result set over 2^32 range.

Project structure:
1. Hash/hash.go  -> Contains the core logic for all the different hashes, same file as project 1.
2. CustomRandom/customrandom.go -> contains the implementation of random number genration logic using unix time, same file as project 1.
3. BloomFilter/bloomfilter.go -> Has the core logic for all three flavors of bloom filter.
4. project1.go -> Contains the main driver functions for all three Bloom filter implementations.(entry point for code - main).

Steps to run
1. Go inside the directory, where project2.go is present.
2. command => go run project2.go

Please check the comment just above the respective functions in main, to change inputs of the functions. Set to demo inputs by default.

Outputs:
All the Outputs are in "output" directory.
1. bloomfilter_output.txt => output for bloom filter.
2. counting_bloomfilter_output.txt => output for counting bloom filter.
3. coded_bloomfilter_output.txt => output for coded bloom filter.

Sample Outputs Acheived are present in "output_recorded" directory for reference, these are for the demo inputs.
