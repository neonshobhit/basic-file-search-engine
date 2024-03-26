# File Search and Sorting Tool

This Go program provides functionality to search for files containing specific words and sorts the files based on their relevance to the search query.

## Usage

1. Clone the repository.

2. Navigate to the directory.

3. Compile and run the program.

## Overview

The program consists of the following main components:

### 1. `readDataFromFiles()`

This function reads data from files located in the `./data/` directory. It constructs an index of words mapped to the files they appear in along with their respective word counts.

### 2. `search(inp string) map[string]File`

This function searches for files containing the words provided in the input string. It returns a map of file names to file objects, where each file object contains the file's ID, name, and ranking based on word occurrence count.

### 3. `sortFiles(inp map[string]File) []File`

This function sorts the files based on their ranking (word occurrence count) in descending order.

### 4. `main()`

The `main()` function is the entry point of the program. It reads data from files, performs a search for the word "Hello", and prints the sorted result.

## File Structure

- `main.go`: Contains the main program logic.
- `./data/`: Directory containing sample data files.

## Dependencies

- The program relies on Go's standard library.

## Sample Output

The program output consists of a list of files sorted by their relevance to the search query "Hello".

## Note

Ensure that the data files are present in the `./data/` directory before running the program. The program expects data files to be plain text files.

Feel free to modify and extend the program to suit your needs!