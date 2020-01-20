package arraylist

/**
 * @Author: eggsy
 * @Description:
 * @File:  iterator
 * @Version: 1.0.0
 * @Date: 2020-01-20 14:22
 */

type Iterator struct {
	list  *List
	index int
}

func (list *List) Iterator() Iterator {
	return Iterator{list: list, index: -1}
}

func (iterator *Iterator) HasNext() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	// 此处完成了判断
	return iterator.list.checkIndex(iterator.index)
}

func (iterator *Iterator) Next() interface{} {
	return iterator.list.elements[iterator.index]
}

func (iterator *Iterator) Remove() {
	iterator.list.Remove(iterator.index)
}
