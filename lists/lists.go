package lists

import (
	"github.com/Eggsyz/pangu/containers"
	"github.com/Eggsyz/pangu/utils"
)

/**
 * @Author: eggsy
 * @Description:
 * @File:  lists
 * @Version: 1.0.0
 * @Date: 2020-01-10 15:46
 */

type List interface {
	Get(index int) (interface{}, bool)
	Remove(index int)
	Add(values ...interface{})
	Containers(values ...interface{}) bool
	Sort(comparator utils.Comparator)
	Swap(index1, index2 int)
	Insert(index int, values ...interface{})
	Set(index int, value interface{})
	containers.Container
}
