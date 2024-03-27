package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type File struct {
	Id      string
	Name    string
	Ranking int
}

type FileIndex struct {
	Id        string
	WordCount int
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

var wordMap map[string]map[string]FileIndex = make(map[string]map[string]FileIndex)
var fileIndex map[string]File = make(map[string]File)
var stopWords map[string]bool = make(map[string]bool)

func fetchStopWords() {
	file, err := os.Open("./stopwords.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		stopWords[word] = true
	}
}

func readDataFromFiles() {
	files, err := os.ReadDir("./data/")
	if err != nil {
		log.Fatal(err)
	}
	for _, fileName := range files {
		if !fileName.IsDir() {
			file, err := os.Open("./data/" + fileName.Name())
			if err != nil {
				panic(err)
			}
			defer file.Close()
			fileIndex[fileName.Name()] = File{Id: fileName.Name(), Name: fileName.Name(), Ranking: 0}
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanWords)

			for scanner.Scan() {
				word := scanner.Text()
				if _, ok := stopWords[word]; ok {
					continue
				}
				_, ok := wordMap[word]
				if !ok {
					wordMap[word] = make(map[string]FileIndex)
				}
				indx := wordMap[word]
				loc, lok := indx[fileName.Name()]
				if lok {
					loc.WordCount++
				} else {
					loc = FileIndex{Id: fileName.Name(), WordCount: 1}
				}
				indx[fileName.Name()] = loc
			}
		}
	}
}

func search(inp string) map[string]File {
	scanner := bufio.NewScanner(strings.NewReader(inp))
	scanner.Split(bufio.ScanWords)

	var result map[string]File = fileIndex
	var temp map[string]File = make(map[string]File)
	for scanner.Scan() {
		word := scanner.Text()
		fileIndx, ok := wordMap[word]
		if ok {
			for k, v := range fileIndx {
				_, filePresent := result[k]
				if filePresent {
					_, fdok := temp[k]
					if !fdok {
						temp[k] = File{Id: k, Name: k, Ranking: 0}
					}
					fileData := temp[k]
					fileData.Ranking += v.WordCount
					temp[k] = fileData
				}
			}
			result = temp
		}
	}
	result = temp

	return result
}

func sortFiles(inp map[string]File) (result []File) {
	var tempResult []File
	for _, v := range inp {
		tempResult = append(tempResult, v)
	}

	sort.Slice(tempResult, func(i, j int) bool {
		return tempResult[i].Ranking > tempResult[j].Ranking
	})
	return tempResult
}
func main() {
	fetchStopWords()
	readDataFromFiles()
	fmt.Println(sortFiles(search("till")))
}
