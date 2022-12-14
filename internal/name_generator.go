package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// NameGenerator 名字生成器
type NameGenerator struct {
	FirstCCS   *CandidateCharacterSet
	MiddleCCS  *CandidateCharacterSet
	LastCCS    *CandidateCharacterSet
	PoemSetMap map[string]*PoemSet
	Evaluator  *Evaluator
}

func NewNameGenerator(
	firstCcs *CandidateCharacterSet,
	middleCcs *CandidateCharacterSet,
	lastCcs *CandidateCharacterSet) *NameGenerator {
	poemSet := make(map[string]*PoemSet)
	nameGenerator := NameGenerator{
		FirstCCS:  firstCcs,
		MiddleCCS: middleCcs,
		LastCCS:   lastCcs,
	}
	nameGenerator.PoemSetMap = poemSet
	nameGenerator.Evaluator = NewEvaluator(poemSet)
	return &nameGenerator
}

func (ng NameGenerator) String() string {
	var result = "FirstName Character Candidates: " + ng.FirstCCS.String() + "\n"
	result += "MiddleName Character Candidates: " + ng.MiddleCCS.String() + "\n"
	result += "LastName Character Candidates: " + ng.LastCCS.String() + "\n"
	return result
}

func (ng NameGenerator) Generate() []*Name {
	var nameList []*Name
	for _, fc := range ng.FirstCCS.CharacterMap {
		for _, mc := range ng.MiddleCCS.CharacterMap {
			for _, lc := range ng.LastCCS.CharacterMap {
				name := NewName(fc, mc, lc)
				if IsValidName(name) {
					nameList = append(nameList, name)
				}
			}
		}
	}
	return nameList
}

func (ng NameGenerator) FindSentencesFitForName(name *Name) []string {
	var fitSentences []string
	for _, p := range ng.PoemSetMap {
		sentences := p.FindSentencesFitForName(name)
		if sentences == nil || len(sentences) <= 0 {
			continue
		}
		fitSentences = append(fitSentences, sentences...)
	}
	return fitSentences
}

func (ng NameGenerator) AddPoemSet(set *PoemSet) {
	setName := set.Name
	ng.PoemSetMap[setName] = set
}

func IsValidName(name *Name) bool {
	if HasDuplicate(name) {
		fmt.Printf("淘汰包含重复的名字：%s\n", name)
		return false
	}
	if BadTuneComposite(name) {
		fmt.Printf("淘汰平仄不合理的名字：%s\n", name)
		return false
	}
	if HasSameTone(name) {
		fmt.Printf("淘汰存在同音字的名字: %s\n", name)
		return false
	}
	if HasSameInitialTone(name) {
		fmt.Printf("淘汰存在相同声母的拗口名字: %s\n", name)
		return false
	}
	if HasSameFinalTone(name) {
		fmt.Printf("淘汰存在相同韵母的拗口名字: %s\n", name)
		return false
	}
	if HasBadFiveElements(name) {
		fmt.Printf("淘汰存在五行相克的名字: %s\n", name)
		return false
	}
	if InBlackList(name) {
		fmt.Printf("淘汰黑名单中的组合名字：%s\n", name)
		return false
	}
	return true
}

func HasDuplicate(name *Name) bool {
	fc := name.FirstCharacter.Character
	mc := name.MiddleCharacter.Character
	lc := name.LastCharacter.Character
	return fc == mc || fc == lc || mc == lc
}

func BadTuneComposite(name *Name) bool {
	ft := name.FirstCharacter.Tune
	mt := name.MiddleCharacter.Tune
	lt := name.LastCharacter.Tune
	tuneTuple := NewTuneTuple3(ft, mt, lt)
	return !PreferredTuneTupleSet.Contains(tuneTuple.String())
}

func HasSameTone(name *Name) bool {
	fc := name.FirstCharacter
	mc := name.MiddleCharacter
	lc := name.LastCharacter
	return IsSameTone(fc, mc) || IsSameTone(mc, lc)
}

func IsSameTone(c1 *Character, c2 *Character) bool {
	it1 := c1.InitialTone
	ft1 := c1.FinalTone
	it2 := c2.InitialTone
	ft2 := c2.FinalTone
	return it1 == it2 && ft1 == ft2
}

func HasSameInitialTone(name *Name) bool {
	fc := name.FirstCharacter
	mc := name.MiddleCharacter
	lc := name.LastCharacter
	return IsSameInitialTone(fc, mc) || IsSameInitialTone(mc, lc)
}

func HasSameFinalTone(name *Name) bool {
	fc := name.FirstCharacter
	mc := name.MiddleCharacter
	lc := name.LastCharacter
	return IsSameFinalTone(fc, mc) || IsSameFinalTone(mc, lc)
}

func IsSameFinalTone(c1 *Character, c2 *Character) bool {
	ft1 := c1.FinalTone
	ft2 := c2.FinalTone
	return ft1 == ft2
}

func IsSameInitialTone(c1 *Character, c2 *Character) bool {
	it1 := c1.InitialTone
	it2 := c2.InitialTone
	return it1 == it2
}

var BlackListNames = NewHashSet()

func init() {
	LoadBlackListNames()
}

func InBlackList(name *Name) bool {
	return BlackListNames.Contains(name.String())
}

func LoadBlackListNames() {
	currentDir, _ := os.Getwd()
	filePath := currentDir + "/assets/blacklist_names.txt"
	LoadBlackListNamesFromFile(filePath)
}

func LoadBlackListNamesFromFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to read file %s, err: %v\n", filePath, err)
		return
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		rawText := fileScanner.Text()
		if strings.HasPrefix(rawText, "#") {
			continue
		}
		if len(rawText) < 1 {
			continue
		}
		BlackListNames.Add(rawText)
	}
}
