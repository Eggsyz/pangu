package containers

import "github.com/Eggsyz/pangu/utils"

/**
 * @Author: eggsy
 * @Description:
 * @File:  containers
 * @Version: 1.0.0
 * @Date: 2020-01-06 19:11
 */

type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
}

func GetSortedValues(container Container, comparator utils.Comparator) []interface{} {
	values := container.Values()
	if len(values) < 2 {
		return values
	}
	utils.Sort(values, comparator)
	return values
}
