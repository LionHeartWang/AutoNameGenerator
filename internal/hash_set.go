package internal

import (
	"bytes"
	"fmt"
)

// HashSet 辅助类
type HashSet struct {
	set map[string] bool
}

func NewHashSet() *HashSet{
	return &HashSet{make(map[string]bool)}
}

func (set *HashSet) Add(i string) bool {
	_, found := set.set[i]
	// 添加元素
	set.set[i] = true
	// 已经有了返回false，否则添加成功
	return !found
}

func (set *HashSet) Contains(i string) bool {
	_, found := set.set[i]
	return found
}

func (set *HashSet) Remove(i string) {
	delete(set.set, i)
}

func (set *HashSet) Print() {
	for k, _ := range set.set {
		fmt.Println(k)
	}
}

func (set *HashSet) Size() int {
	return len(set.set)
}

func (set *HashSet) FormatToString() string {
	var buf bytes.Buffer
	for k, _ := range set.set {
		buf.WriteString(k+"\n")
	}
	return buf.String()
}