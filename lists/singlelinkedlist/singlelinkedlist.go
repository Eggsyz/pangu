package linkedlist

import "github.com/Eggsyz/pangu/utils"

/**
 * @Author: eggsy
 * @Description: 单向链表
 * @File:  singlelinkedlist
 * @Version: 1.0.0
 * @Date: 2020-01-17 19:11
 */

// 保存head、tail以及size
type List struct {
	head *element
	tail *element
	size int
}

// 元素包含当前值以及下一个元素的地址
type element struct {
	value interface{}
	next  *element
}

func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values)
	}
	return list
}

func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		newElement := &element{value: value}
		if list.size == 0 {
			list.head = newElement
			list.tail = newElement
		} else {
			list.tail.next = newElement
			list.tail = list.tail.next
		}
		list.size++
	}
}

func (list *List) Get(index int) (interface{}, bool) {
	if list.checkIndex(index) {
		tmp := list.head
		for i := 0; i < index; i++ {
			tmp = tmp.next
		}
		return tmp.value, true
	} else {
		return nil, false
	}
}

func (list *List) checkIndex(index int) bool {
	return 0 <= index && index < list.size
}

func (list *List) Remove(index int) {
	if list.checkIndex(index) {
		if list.size == 1 {
			list.Clear()
			return
		}
		current := list.head
		var before *element
		for i := 0; i < index; i, current = i+1, current.next {
			before = current
		}
		if current == list.head {
			list.head = list.head.next
		}
		if current == list.tail {
			list.tail = before
		}
		if before != nil {
			before.next = current.next
		}
		current = nil
		list.size--
	}
}

func (list *List) Containers(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}
	if list.size == 0 {
		return false
	}
	for _, value := range values {
		var flag bool
		for element := list.head; element != nil; element = element.next {
			if element.value == value {
				flag = true
				break
			}
		}
		if flag == false {
			return false
		}
	}
	return true
}

func (list *List) Sort(comparator utils.Comparator) {
	values := list.Values()
	utils.Sort(values, comparator)
	list.Clear()
	list.Add(values...)
}

func (list *List) Swap(index1, index2 int) {
	var element1, element2 *element
	if list.checkIndex(index1) && list.checkIndex(index2) {
		for i, cur := 0, list.head; cur != nil; i, cur = i+1, cur.next {
			switch i {
			case index1:
				element1 = cur
			case index2:
				element2 = cur
			}
		}
		element1.value, element2.value = element2.value, element1.value
	}
}

func (list *List) Insert(index int, values ...interface{}) {
	if list.checkIndex(index) {
		list.size += len(values)
		current := list.head
		var before *element
		for i := 0; i < index; i, current = i+1, current.next {
			before = current
		}
		if current == list.head {
			oldNextElement := list.head
			for i, value := range values {
				newElement := &element{value: value}
				if i == 0 {
					list.head = newElement
				} else {
					// 临时保存
					before.next = newElement
				}
				before = newElement
			}
			before.next = oldNextElement
		} else {
			oldNextElement := before.next
			for _, value := range values {
				newElement := &element{value: value}
				before.next = newElement
				before = before.next
			}
			before.next = oldNextElement
		}
	} else {
		if index == list.size {
			list.Add(values)
		}
	}
}

func (list *List) Set(index int, value interface{}) {
	if list.checkIndex(index) {
		cur := list.head
		for i := 0; i < index; i++ {
			cur = cur.next
		}
		cur.value = value
	} else {
		// append
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
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *List) Values() []interface{} {
	var values = make([]interface{}, list.size, list.size)
	for i, element := 0, list.head; element != nil; i, element = i+1, element.next {
		values[i] = element.value
	}
	return values
}
