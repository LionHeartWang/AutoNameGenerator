package main

import (
	"AutoNameGenerator/internal"
	"fmt"
	"os"
)

func main() {
	firstCcs, err := LoadCandidateCharacterSet(
		"first", "first_character_candidate_set.txt")
	if err != nil {
		fmt.Println("Failed to load first character candidate set!")
		return
	}
	middleCcs, err := LoadCandidateCharacterSet(
		"middle", "middle_character_candidate_set.txt")
	if err != nil {
		fmt.Println("Failed to load middle character candidate set!")
		return
	}
	lastCcs, err := LoadCandidateCharacterSet(
		"last", "last_character_candidate_set.txt")
	if err != nil {
		fmt.Println("Failed to load last character candidate set!")
		return
	}
	nameGenerator := internal.NewNameGenerator(firstCcs, middleCcs, lastCcs)
	nameList := nameGenerator.Generate()
	fmt.Println("自动起名结果：")
	for idx, name := range nameList {
		fmt.Printf("No.%d %s\n", idx, name.Explain())
	}
}

func LoadCandidateCharacterSet(
	name string, fileName string) (*internal.CandidateCharacterSet, error) {
	currentDir, _ := os.Getwd()
	filePath := currentDir + "/assets/" + fileName
	return internal.LoadCandidateCharacterSetFromFile(name, filePath)
}
