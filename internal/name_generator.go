package internal

import "fmt"

/**
 * 名字生成器。
 */
type NameGenerator struct {
	FirstCCS  *CandidateCharacterSet
	MiddleCCS *CandidateCharacterSet
	LastCCS   *CandidateCharacterSet
}

func NewNameGenerator(
	firstCcs *CandidateCharacterSet,
	middleCcs *CandidateCharacterSet,
	lastCcs *CandidateCharacterSet) *NameGenerator {
	return &NameGenerator{
		FirstCCS:  firstCcs,
		MiddleCCS: middleCcs,
		LastCCS:   lastCcs,
	}
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

func IsValidName(name *Name) bool {
	if HasDuplicate(name) {
		fmt.Printf("淘汰包含重复的名字：%s\n", name)
		return false
	}
	if BadTuneComposite(name) {
		fmt.Printf("淘汰平仄不合理的名字：%s\n", name)
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
