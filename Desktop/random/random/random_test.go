package random

import "testing"

func hasDup(list []int) bool {
	dup_map := make(map[int]int)

	for _, item := range list {
		_, exist := dup_map[item]

		if exist {
			return true
		}
	}

	return false
}

func TestRandomness(t *testing.T) {
	rlist := MiddleSquare(304393)
	rlist2 := LinearCongruentialGenerator(823312, 272142, 726462, 471242)
	rlist3 := Xorshift(912832)

	hasDupList := make([]bool, 3)
	hasDupList[0] = hasDup(rlist)
	hasDupList[1] = hasDup(rlist2)
	hasDupList[2] = hasDup(rlist3)

	for _, item := range hasDupList {
		if item {
			t.Fatalf("Test failed: Random list has duplicate entry")
		}
	}
}
