package linkedlist

import (
	"fmt"
	"github.com/Eggsyz/pangu/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author: eggsy
 * @Description:
 * @File:  singlelinkedlist_test
 * @Version: 1.0.0
 * @Date: 2020-01-20 17:19
 */

func TestList_New(t *testing.T) {
	list := New()
	// 测试add
	list.Add(0, 1, 2, 3)
	for i := 0; i < 4; i++ {
		// 顺带测试get
		v, _ := list.Get(i)
		assert.True(t, v == i, "excepting %v but found %v", i, v)
	}
	// 测试containers
	res := list.Containers(0, 1)
	assert.True(t, res == true, "excepting true but found %v", res)
	res = list.Containers(0, 9)
	assert.True(t, res == false, "excepting false but found %v", res)
	// 测试set
	list.Set(0, 4)
	v, _ := list.Get(0)
	assert.True(t, v == 4, "excepting 4 but found %v", v)
	// 测试sort
	list.Sort(utils.IntComparator)
	values := list.Values()
	for i := 0; i < len(values); i++ {
		assert.True(t, values[i] == i+1, "excepting %v but found %v", i+1, values[i])
	}
	// 测试swap
	v0, _ := list.Get(0)
	v2, _ := list.Get(2)
	list.Swap(0, 2)
	v_0, _ := list.Get(0)
	v_2, _ := list.Get(2)
	assert.True(t, v0 == v_2 && v2 == v_0, "excepting true but found %v", v0 == v_2 && v2 == v_0)
	// 测试insert
	list.Insert(1, 5)
	v, _ = list.Get(1)
	assert.True(t, v == 5, "excepting 5 but found %v", v)
}

func TestList_Iterator(t *testing.T) {
	list := New()
	list.Add(0, 1, 2, 3)
	iterator := list.Iterator()
	var i = 0
	for iterator.HasNext() {
		v := iterator.Next()
		assert.True(t, v == i, "excepting i but found %v", v)
		i++
		if v == 3 {
			iterator.Remove()
			fmt.Println(list.Values())
		}
	}
}
