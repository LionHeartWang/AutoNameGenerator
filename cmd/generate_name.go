package main

import (
	"AutoNameGenerator/internal"
	"fmt"
	"os"
)

func main() {
	poemSets, err := InitPoemSets()
	if err != nil {
		return
	}

	firstCcs, middleCcs, lastCcs, err := InitCharacterSet()
	if err != nil {
		return
	}

	nameGenerator := internal.NewNameGenerator(firstCcs, middleCcs, lastCcs)
	for _, ps := range poemSets {
		nameGenerator.AddPoemSet(ps)
	}

	nameList := nameGenerator.Generate()
	fmt.Println("自动起名结果：")
	for idx, name := range nameList {
		fmt.Printf("No.%d %s\n", idx, name.Explain(poemSets))
	}
}

func InitPoemSets() ([]*internal.PoemSet, error) {
	tangPoems300, err := LoadPoemsSet(
		"唐诗", "tang_poems_300.txt")
	var poemSets []*internal.PoemSet
	poemSets = append(poemSets, tangPoems300)
	return poemSets, err
}

func InitCharacterSet() (
	*internal.CandidateCharacterSet,
	*internal.CandidateCharacterSet,
	*internal.CandidateCharacterSet,
	error) {
	firstCcs, err := LoadCandidateCharacterSet(
		"first", "first_character_candidate_set.txt")
	if err != nil {
		fmt.Println("Failed to load first character candidate set!")
		return nil, nil, nil, err
	}
	middleCcs, err := LoadCandidateCharacterSet(
		"middle", "middle_character_candidate_set.txt")
	if err != nil {
		fmt.Println("Failed to load middle character candidate set!")
		return nil, nil, nil, err
	}
	lastCcs, err := LoadCandidateCharacterSet(
		"last", "last_character_candidate_set.txt")
	if err != nil {
		fmt.Println("Failed to load last character candidate set!")
		return nil, nil, nil, err
	}

	return firstCcs, middleCcs, lastCcs, err
}

func LoadCandidateCharacterSet(
	name string, fileName string) (*internal.CandidateCharacterSet, error) {
	currentDir, _ := os.Getwd()
	filePath := currentDir + "/assets/" + fileName
	return internal.LoadCandidateCharacterSetFromFile(name, filePath)
}

func LoadPoemsSet(name string, fileName string) (*internal.PoemSet, error) {
	currentDir, _ := os.Getwd()
	filePath := currentDir + "/assets/" + fileName
	return internal.LoadPoemSetFromFile(name, filePath)
}
