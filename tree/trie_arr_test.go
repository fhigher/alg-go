package tree

import (
	"fmt"
	"testing"
)

func TestTrieArr(t *testing.T) {
	names := []string{"jerry", "david", "lee"}
	needle := []string{"lee", "jerry", "joe"}

	for _, name := range names {
		Insert(name)
	}

	for _, name := range needle {
		if Find(name) {
			fmt.Printf("find it: %s\n", name)
		} else {
			fmt.Printf("cannot find it: %s\n", name)
		}
	}

	fmt.Println(ID)
	//fmt.Println(Vector)
	fmt.Println(Flag)
}

func TestSub(t *testing.T) {
	fmt.Println('l' - 'a')
}
