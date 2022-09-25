package internal

/**
 * 名字定义。
 */
type Name struct {
	FirstCharacter  *CandidateCharacter
	MiddleCharacter *CandidateCharacter
	LastCharacter   *CandidateCharacter
}

func NewName(
	fc *CandidateCharacter, mc *CandidateCharacter, lc *CandidateCharacter) *Name {
	return &Name{
		FirstCharacter:  fc,
		MiddleCharacter: mc,
		LastCharacter:   lc,
	}
}

func (ng Name) String() string {
	return ng.FirstCharacter.Character + ng.MiddleCharacter.Character + ng.LastCharacter.Character
}