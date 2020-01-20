package utils

import "testing"

/**
 * @Author: eggsy
 * @Description:
 * @File:  sort_test
 * @Version: 1.0.0
 * @Date: 2019-10-22 11:38
 */

func TestSortInts(t *testing.T) {
	ints := []interface{}{}
	ints = append(ints, 4)
	ints = append(ints, 1)
	ints = append(ints, 2)
	ints = append(ints, 3)

	Sort(ints, IntComparator)

	for i := 1; i < len(ints); i++ {
		if ints[i-1].(int) > ints[i].(int) {
			t.Errorf("Not sorted!")
		}
	}
}

func TestSortStrings(t *testing.T) {
	strings := []interface{}{}
	strings = append(strings, "a")
	strings = append(strings, "abc")
	strings = append(strings, "ab")
	strings = append(strings, "abcd")

	Sort(strings, StringComparator)

	for i := 1; i < len(strings); i++ {
		if strings[i-1].(string) > strings[i].(string) {
			t.Errorf("Not sorted!")
		}
	}
}

func TestSortStruct(t *testing.T) {
	type Custom struct {
		id int
	}
	comparatorById := func(a, b interface{}) int {
		aAsserted := a.(Custom)
		bAsserted := b.(Custom)
		switch {
		case aAsserted.id > bAsserted.id:
			return 1
		case aAsserted.id < bAsserted.id:
			return -1
		default:
			return 0
		}
	}
	customs := []interface{}{}
	customs = append(customs, Custom{1})
	customs = append(customs, Custom{1})
	customs = append(customs, Custom{5})
	customs = append(customs, Custom{2})

	Sort(customs, comparatorById)

	for i := 1; i < len(customs); i++ {
		if customs[i-1].(Custom).id > customs[i].(Custom).id {
			t.Errorf("Not sorted!")
		}
	}
}
