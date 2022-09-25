package internal


/**
 * 名字定义。
 */
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