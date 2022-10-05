package internal

// Evaluator 名字评分器
type Evaluator struct {
	PreferredCharacters *HashSet
	PoemSet             map[string]*PoemSet
}

func NewEvaluator(poemSetMap map[string]*PoemSet) *Evaluator {
	preferredCharacterSet := DefaultPreferredCharacters
	return &Evaluator{
		PreferredCharacters: preferredCharacterSet,
		PoemSet:             poemSetMap,
	}
}

var DefaultPreferredCharacters = NewHashSet()

func (evaluator Evaluator) Evaluate(name *Name) int {
	score := 0
	score += evaluator.CalFiveElementsScore(name)
	score += evaluator.CalPoemScore(name)
	score += evaluator.CalPreferredCharacterScore(name)
	return score
}

func (evaluator Evaluator) CalPoemScore(name *Name) int {
	score := 0
	for _, p := range evaluator.PoemSet {
		sentences := p.FindSentencesFitForName(name)
		if sentences != nil && len(sentences) > 0 {
			score += 1
			break
		}
	}
	return score
}

func (evaluator Evaluator) CalFiveElementsScore(name *Name) int {
	score := 0
	fc := name.FirstCharacter
	mc := name.MiddleCharacter
	lc := name.LastCharacter
	score += CalCharacterFEScore(fc, mc)
	score += CalCharacterFEScore(mc, lc)
	return score
}

func (evaluator Evaluator) CalPreferredCharacterScore(name *Name) int {
	mc := name.MiddleCharacter.Character
	lc := name.LastCharacter.Character
	if evaluator.PreferredCharacters.Contains(mc) || evaluator.PreferredCharacters.Contains(lc) {
		return 1
	}
	return 0
}

func CalCharacterFEScore(c1 *Character, c2 *Character) int {
	fe1 := c1.FiveElements
	fe2 := c2.FiveElements
	if IsGoodFiveElements(fe1, fe2) || IsGoodFiveElements(fe2, fe1) {
		return 1
	}
	if IsBadFiveElements(fe1, fe2) || IsBadFiveElements(fe2, fe1) {
		return -1
	}
	return 0
}

func init() {
	var characters = []string{
		"川",
		"枫",
	}
	for _, v := range characters {
		DefaultPreferredCharacters.Add(v)
	}
}
