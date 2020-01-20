package utils

import (
	"testing"
	"time"
)

/**
 * @Author: eggsy
 * @Description:
 * @File:  comparator_test
 * @Version: 1.0.0
 * @Date: 2019-10-22 11:09
 */

func TestIntComparator(t *testing.T) {
	tests := [][]interface{}{
		{1, 1, 0},
		{1, 0, 1},
		{1, 2, -1},
		{-1, -2, 1},
		{-2, -1, -1},
		{-2, 0, -1},
	}
	for _, test := range tests {
		actual := IntComparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestStringComparator(t *testing.T) {
	tests := [][]interface{}{
		{"abcd", "abcd", 0},
		{"abcd", "", 1},
		{"abcd", "bcde", -1},
		{"", "abcd", -1},
		{"bcde", "abcd", 1},
	}
	for _, test := range tests {
		actual := StringComparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestTimeComparator(t *testing.T) {
	now := time.Now()
	tests := [][]interface{}{
		{now, now, 0},
		{now.Add(time.Duration(time.Second * 10)), now, 1},
		{now, now.Add(time.Hour), -1},
	}
	for _, test := range tests {
		actual := TimeComparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestCustomComparator(t *testing.T) {
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
	tests := [][]interface{}{
		{Custom{id: 0}, Custom{id: 0}, 0},
		{Custom{id: -1}, Custom{id: 0}, -1},
		{Custom{id: 1}, Custom{id: 0}, 1},
	}
	for _, test := range tests {
		actual := comparatorById(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}
