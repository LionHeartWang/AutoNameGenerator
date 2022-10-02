package internal

// Name 名字定义
type Name struct {
	FirstCharacter  *Character
	MiddleCharacter *Character
	LastCharacter   *Character
}

func NewName(
	fc *Character, mc *Character, lc *Character) *Name {
	return &Name{
		FirstCharacter:  fc,
		MiddleCharacter: mc,
		LastCharacter:   lc,
	}
}

func (name Name) String() string {
	return name.FirstCharacter.Character + name.MiddleCharacter.Character + name.LastCharacter.Character
}

func (name Name) Explain() string {
	var result = name.String() + ":"
	tuneComposite := NewTuneTuple3(
		name.FirstCharacter.Tune, name.MiddleCharacter.Tune, name.LastCharacter.Tune)
	result += tuneComposite.Explain() + ", 五行属"
	result += name.FirstCharacter.FiveElements.String()
	result += name.MiddleCharacter.FiveElements.String()
	result += name.LastCharacter.FiveElements.String()
	if HasGoodFiveElements(&name) {
		result += ", 五行相生("
		isFirst := true
		if IsGoodFiveElements(
			name.FirstCharacter.FiveElements, name.MiddleCharacter.FiveElements) {
			result += name.FirstCharacter.FiveElements.String() + "生" + name.MiddleCharacter.FiveElements.String()
			isFirst = false
		}
		if IsGoodFiveElements(
			name.MiddleCharacter.FiveElements, name.LastCharacter.FiveElements) {
			if !isFirst {
				result += ", "
			}
			result += name.MiddleCharacter.FiveElements.String() + "生" + name.LastCharacter.FiveElements.String()
		}
		result += ")"
	}
	result += ", 寓意为"
	meaning1 := name.MiddleCharacter.MeaningList
	if len(meaning1) > 0 {
		result += meaning1[0]
	}
	meaning2 := name.LastCharacter.MeaningList
	if len(meaning2) > 0 {
		result += "、" + meaning2[0]
	}

	result += "。"

	return result
}
