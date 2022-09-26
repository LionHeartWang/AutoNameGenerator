package internal

import (
	"errors"
	"fmt"
	"strings"
)

// 声调定义
type Tune int

const (
	TuneFirst  = 0
	TuneSecond = 1
	TuneThird  = 2
	TuneFourth = 3
)

func (t Tune) String() string {
	return [...]string{"一声", "二声", "三声", "四声"}[t]
}

func Int2Tune(id int) (Tune, error) {
	switch id {
	case 1:
		return TuneFirst, nil
	case 2:
		return TuneSecond, nil
	case 3:
		return TuneThird, nil
	case 4:
		return TuneFourth, nil
	default:
		message := fmt.Sprintf("非法的声调取值: %d", id)
		return -1, errors.New(message)
	}
}

// 五行定义
type FiveElements int

const (
	GOLD  = 0 // 金
	WOOD  = 1 // 木
	WATER = 2 // 水
	FIRE  = 3 // 火
	SOIL  = 4 // 土
)

func (fe FiveElements) String() string {
	return [...]string{"金", "木", "水", "火", "土"}[fe]
}

func Str2FiveElements(str string) (FiveElements, error) {
	switch str {
	case "金":
		return GOLD, nil
	case "木":
		return WOOD, nil
	case "水":
		return WATER, nil
	case "火":
		return FIRE, nil
	case "土":
		return SOIL, nil
	default:
		message := fmt.Sprintf("%s不属于五行元素。", str)
		return -1, errors.New(message)
	}
}

/**
 * 候选字。
 */
type Character struct {
	// 汉字
	Character string

	// 声调
	Tune Tune

	// 五行
	FiveElements FiveElements

	// 声母
	InitialTone Tone

	// 韵母
	FinalTone Tone

	// 寓意
	MeaningList []string
}

func (cc Character) String() string {
	meaning := strings.Join(cc.MeaningList, "、")
	return fmt.Sprintf("%s: %s, %s属性, 寓意为%s\n", cc.Character, cc.Tune, cc.FiveElements, meaning)
}

func NewCharacter(
	name string, tuneId int, initialToneStr string, finalToneStr string,
	fiveElementsStr string, meaning string) (*Character, error) {
	tune, err := Int2Tune(tuneId)
	if err != nil {
		return nil, err
	}
	fiveElements, err := Str2FiveElements(fiveElementsStr)
	initialTone := Str2Tone(initialToneStr)
	finalTone := Str2Tone(finalToneStr)
	meaningList := strings.Split(meaning, "|")
	return &Character{
		Character:    name,
		Tune:         tune,
		FiveElements: fiveElements,
		InitialTone:  initialTone,
		FinalTone:    finalTone,
		MeaningList:  meaningList,
	}, nil
}
