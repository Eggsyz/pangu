package utils

import "sort"

/**
 * @Author: eggsy
 * @Description:
 * @File:  sort
 * @Version: 1.0.0
 * @Date: 2019-10-22 11:32
 */

func Sort(value []interface{}, comparator Comparator) {
	sort.Sort(sortable{value: value, comparator: comparator})
}

type sortable struct {
	value      []interface{}
	comparator Comparator
}

func (s sortable) Len() int {
	return len(s.value)
}

func (s sortable) Less(i, j int) bool {
	return s.comparator(s.value[i], s.value[j]) < 0
}
func (s sortable) Swap(i, j int) {
	s.value[i], s.value[j] = s.value[j], s.value[i]
}
