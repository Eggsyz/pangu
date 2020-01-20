package linkedlist

/**
 * @Author: eggsy
 * @Description:
 * @File:  iterator
 * @Version: 1.0.0
 * @Date: 2020-01-20 16:42
 */
type Iterator struct {
	list    *List
	index   int
	element *element
}

func (list *List) Iterator() Iterator {
	return Iterator{list: list, index: -1, element: nil}
}

// 功能：判断是否存在下一个，并且指向下一个element
func (iterator *Iterator) HasNext() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	if !iterator.list.checkIndex(iterator.index) {
		iterator.element = nil
		return false
	}
	// 转到下一个element
	if iterator.index == 0 {
		iterator.element = iterator.list.head
	} else {
		iterator.element = iterator.element.next
	}
	return true
}

func (iterator *Iterator) Next() interface{} {
	return iterator.element.value
}

func (iterator *Iterator) Remove() {
	iterator.list.Remove(iterator.index)
}
