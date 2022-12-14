package internal

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Poem 诗歌定义
type Poem struct {
	Title     string
	Author    string
	Sentences *[]string
}

func (poem Poem) Add(sentence string) {
	*poem.Sentences = append(*poem.Sentences, sentence)
}

func (poem Poem) FindSentencesFitForName(name *Name) []string {
	mc := name.MiddleCharacter
	lc := name.LastCharacter
	var fitSentences []string
	for _, s := range *poem.Sentences {
		if mc != nil && !strings.Contains(s, mc.Character) {
			continue
		}
		if lc != nil && !strings.Contains(s, lc.Character) {
			continue
		}
		fitSentences = append(fitSentences, s)
	}
	return fitSentences
}

func (poem Poem) String() string {
	var result = "《" + poem.Title + "》 " + poem.Author

	for _, s := range *poem.Sentences {
		result += s
		result += "\n"
	}
	return result
}

func (poem Poem) Print() {
	fmt.Println("《" + poem.Title + "》 " + poem.Author)
	for _, s := range *poem.Sentences {
		fmt.Println(s)
	}
}

func NewEmptyPoem(title string, author string) *Poem {
	var sentences []string
	return &Poem{
		Title:     title,
		Author:    author,
		Sentences: &sentences,
	}
}

func NewPoem(title string, author string, sentences []string) *Poem {
	return &Poem{
		Title:     title,
		Author:    author,
		Sentences: &sentences,
	}
}

// PoemSet 诗集定义
type PoemSet struct {
	Name    string
	PoemMap map[string]*Poem
}

func (poemSet PoemSet) Add(poem *Poem) {
	if poem != nil {
		title := poem.Title
		poemSet.PoemMap[title] = poem
	}
}

func (poemSet PoemSet) FindSentencesFitForName(name *Name) []string {
	var fitSentences []string
	for _, p := range poemSet.PoemMap {
		sentences := p.FindSentencesFitForName(name)
		if sentences == nil || len(sentences) <= 0 {
			continue
		}
		fitSentences = append(fitSentences, sentences...)
	}
	return fitSentences
}

func (poemSet PoemSet) Remove(title string) {
	delete(poemSet.PoemMap, title)
}

func (poemSet PoemSet) String() string {
	var result = poemSet.Name
	i := 0
	for k, _ := range poemSet.PoemMap {
		result += fmt.Sprintf("No.%d %s\n", i, k)
		i++
	}
	return result
}

func (poemSet PoemSet) Print() {
	fmt.Println("《" + poemSet.Name + "》")
	for _, p := range poemSet.PoemMap {
		p.Print()
	}
}

func NewPoemSet(name string) *PoemSet {
	poemMap := make(map[string]*Poem)
	return &PoemSet{
		Name:    name,
		PoemMap: poemMap,
	}
}

func LoadPoemSetFromFile(name string, filePath string) (*PoemSet, error) {
	poemSet := NewPoemSet(name)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to read file %s, err: %v\n", filePath, err)
		return poemSet, err
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	var currentPoem *Poem = nil
	for fileScanner.Scan() {
		rawText := fileScanner.Text()
		if strings.HasPrefix(rawText, "#") {
			continue
		}
		PATTERN := `(\d+)(.+)：(.+)`
		r := regexp.MustCompile(PATTERN)
		matches := r.FindStringSubmatch(rawText)
		if matches != nil && len(matches) >= 3 {
			// 匹配到标题行
			author := matches[2]
			title := matches[3]
			currentPoem = NewEmptyPoem(title, author)
			poemSet.Add(currentPoem)
		} else {
			// 匹配到正文行
			if len(rawText) > 0 {
				currentPoem.Add(rawText)
			}
		}
	}

	if err := fileScanner.Err(); err != nil {
		fmt.Printf("Error while reading file: %s", err)
		return poemSet, err
	}

	return poemSet, err
}
