package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * 候选集。
 */
type CandidateCharacterSet struct {
	Name         string
	CharacterMap map[string]*Character
}

func (ccs CandidateCharacterSet) Add(cc *Character) {
	if cc != nil {
		character := cc.Character
		ccs.CharacterMap[character] = cc
	}
}

func (ccs CandidateCharacterSet) Remove(character string) {
	delete(ccs.CharacterMap, character)
}

func (ccs CandidateCharacterSet) String() string {
	var result = ccs.Name
	for _, v := range ccs.CharacterMap {
		result += v.String()
		result += "\n"
	}
	return result
}

func NewCandidateCharacterSet(name string) *CandidateCharacterSet {
	characterMap := make(map[string]*Character)
	return &CandidateCharacterSet{
		Name:         name,
		CharacterMap: characterMap,
	}
}

func LoadCandidateCharacterSetFromFile(name string, filePath string) (*CandidateCharacterSet, error) {
	ccs := NewCandidateCharacterSet(name)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to read file %s, err: %v\n", filePath, err)
		return ccs, err
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		rawText := fileScanner.Text()
		fields := strings.Split(rawText, ",")
		if len(fields) != 6 {
			continue
		}
		tune, idErr := strconv.Atoi(fields[1])
		if idErr != nil {
			return ccs, idErr
		}

		cs, csErr := NewCharacter(fields[0], tune, fields[3], fields[4], fields[2], fields[5])
		if csErr != nil {
			return ccs, csErr
		}

		ccs.Add(cs)
	}

	if err := fileScanner.Err(); err != nil {
		fmt.Printf("Error while reading file: %s", err)
		return ccs, err
	}

	return ccs, err
}
