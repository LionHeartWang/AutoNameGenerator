package internal

import (
	"fmt"
	"testing"
)

func TestPutGetHashSet(t *testing.T) {
	hashSet := NewHashSet()
	tuple1, _ := NewTuneTuple3FromNum(1,2,3)
	hashSet.Put(tuple1)
	tuple2, _ := NewTuneTuple3FromNum(4,5,6)
	hashSet.Put(tuple2)
	tuple3, _ := NewTuneTuple3FromNum(1,2,3)
	hashSet.Put(tuple3)

	fmt.Printf("Hashset size: %d\n", hashSet.Len)

	tuple4, _ := NewTuneTuple3FromNum(7,8,9)
	if hashSet.IsExist(tuple4) {
		fmt.Println("Hashset contains tuple4.")
	} else {
		fmt.Println("Hashset doesn't contain tuple4.")
	}

	tuple5, _ := NewTuneTuple3FromNum(4,5,6)
	if hashSet.IsExist(tuple5) {
		fmt.Println("Hashset contains tuple5.")
	} else {
		fmt.Println("Hashset doesn't contain tuple5.")
	}
}
