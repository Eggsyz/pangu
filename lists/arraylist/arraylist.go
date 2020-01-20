package arraylist

import "github.com/Eggsyz/pangu/utils"

/**
 * @Author: eggsy
 * @Description:
 * @File:  arraylist
 * @Version: 1.0.0
 * @Date: 2020-01-17 19:05
 */

type List struct {
	elements []interface{}
	size     int
}

func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values)
	}
	return list
}

func (list *List) Add(values ...interface{}) {
	list.growList(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

func (list *List) Get(index int) (interface{}, bool) {
	if list.checkIndex(index) {
		return list.elements[index], true
	} else {
		return nil, false
	}
}

func (list *List) Remove(index int) {
	if list.checkIndex(index) {
		list.elements[index] = nil
		copy(list.elements[index:], list.elements[index+1:])
		list.size--
	} else {
		return
	}
}

// 暂时没有发现很好的办法，只能双重for循环
func (list *List) Containers(values ...interface{}) bool {
	for _, value := range values {
		var flag bool
		for _, element := range list.elements {
			if value == element {
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}

func (list *List) Sort(comparator utils.Comparator) {
	if list.size < 2 {
		return
	}
	utils.Sort(list.elements[:list.size], comparator)
}

func (list *List) Swap(index1, index2 int) {
	if list.checkIndex(index1) && list.checkIndex(index2) {
		list.elements[index1], list.elements[index2] = list.elements[index2], list.elements[index1]
	}
}

func (list *List) Insert(index int, values ...interface{}) {
	if !list.checkIndex(index) {
		if index == list.size {
			list.Add(values)
		}
		return
	}
	var length = len(values)
	// 1. 拓展
	list.growList(length)
	// 2. 腾出位置
	list.size += length
	copy(list.elements[index+length:], list.elements[index:list.size-length])
	// 3. 插入values
	copy(list.elements[index:], values)
}

func (list *List) Set(index int, value interface{}) {
	if list.checkIndex(index) {
		list.elements[index] = value
	} else {
		if index == list.size {
			list.Add(value)
		}
	}
}

func (list *List) Empty() bool {
	return list.size == 0
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

func (list *List) Values() []interface{} {
	return list.elements
}

func (list *List) IndexOf(value interface{}) int {
	if list.size == 0 {
		return -1
	}
	for index, element := range list.elements {
		if element == value {
			return index
		}
	}
	return -1
}

func (list *List) roundupsize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

// 判断index是否在范围内
func (list *List) checkIndex(index int) bool {
	return 0 <= index && index < list.size
}

// 实现arraylist扩容，当长度在1024以下时，扩容为原来的2倍，超过1024，扩容为原来的1.25
// n代表当前操作需要添加n个element
func (list *List) growList(n int) {
	newCap := cap(list.elements)
	requireSize := list.size + n
	if requireSize > newCap { //判断是否需要扩容
		doubleCap := newCap + newCap
		if requireSize > doubleCap {
			newCap = requireSize
		} else {
			if len(list.elements) < 1024 { // 长度小于1024
				newCap = doubleCap
			} else {
				for 0 < newCap && newCap < requireSize { //扩容为原来的1.25不满足则继续扩容
					newCap += newCap / 4
				}
				if newCap <= 0 {
					newCap = requireSize
				}
			}
		}
		// 执行更新操作
		list.roundupsize(newCap)
	}
}
