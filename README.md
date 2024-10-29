# WC
This repository provides a solution for [Coding Challenge #01](https://codingchallenges.fyi/challenges/challenge-wc), which involves creating a custom implementation of the wc (word count) tool. The project replicates core functionalities of wc, offering counts of lines, words, and characters from text inputs.

## Build and Run

To build and run the tool, use the following commands:

```sh
# Build the project
make build

# Run the wc tool on a file to display line, word, and character counts
./bin/wc path/to/file.txt

# Use flags for specific counts, e.g., word count only
./bin/wc -w path/to/file.txt

# Display line count only
./bin/wc -l path/to/file.txt

# Display character count only
./bin/wc -c path/to/file.txt
```

## Test

To run tests for the tool, use the following command:

```sh
# Run tests
make test
```
