package internal

import (
	"fmt"
	"testing"
)


func TestPutGetHashSet(t *testing.T) {
	hashSet := NewHashSet()
	tuple1, _ := NewTuneTuple3FromNum(1,2,3)
	_ = hashSet.Add(tuple1.String())
	tuple2, _ := NewTuneTuple3FromNum(1,2,4)
	_ = hashSet.Add(tuple2.String())
	tuple3, _ := NewTuneTuple3FromNum(1,2,3)
	_ = hashSet.Add(tuple3.String())
	fmt.Printf("Hashset size: %d\n", hashSet.Size())

	tuple4, _ := NewTuneTuple3FromNum(2,1,3)
	if hashSet.Contains(tuple4.String()) {
		fmt.Println("Hashset contains tuple4.")
	} else {
		fmt.Println("Hashset doesn't contain tuple4.")
	}

	tuple5, _ := NewTuneTuple3FromNum(1,2,4)
	if hashSet.Contains(tuple5.String()) {
		fmt.Println("Hashset contains tuple5.")
	} else {
		fmt.Println("Hashset doesn't contain tuple5.")
	}
}
