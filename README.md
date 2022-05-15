## Rolling Hash Algorithm File Differ
This repo contains the implementation for Rolling Hash Algorithm File Differ. This differ takes two files and using the 
rolling hash algorithm compares the two files and returns two slices of chunks of the two files that match and differ.

## Running the application
To run the application run `make run` command. This will calculate the delta between the two files and will print out the 
delta stats.

## Unit Tests
To run all unit tests run `make tests` command. This will run all the tests and will show the coverage as well.