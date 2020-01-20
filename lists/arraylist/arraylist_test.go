package arraylist

import (
	"github.com/Eggsyz/pangu/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author: eggsy
 * @Description:
 * @File:  arraylist_test
 * @Version: 1.0.0
 * @Date: 2020-01-20 11:14
 */

func TestList_New(t *testing.T) {
	list := New()
	// 测试长度Size()
	assert.True(t, list.Size() == 0, "excepting 0 but found %v", list.Size())
	// 测试Empty
	assert.True(t, list.Empty() == true, "excepting true but found %v", list.Empty())
	// 测试Get
	_, ok := list.Get(0)
	assert.True(t, ok == false, "excepting false but found %v", ok)
	// 测试Add
	list.Add(0, 2, 1, 3)
	assert.True(t, list.Empty() == false, "excepting false but found %v", list.Empty())
	value, ok := list.Get(1)
	assert.True(t, ok == true, "excepting true but found %v", ok)
	assert.True(t, value == 2, "excepting 2 but found %v", value)
	// 测试indexOf
	index := list.IndexOf(3)
	assert.True(t, index == 3, "excepting 3 but found %v", index)
	index = list.IndexOf(9)
	assert.True(t, index == -1, "excepting -1 but found %v", index)
	//测试sort
	list.Sort(utils.IntComparator)
	// 测试values
	values := list.Values()
	for i := 0; i < len(values); i++ {
		assert.True(t, values[i] == i, "excepting %v but found %v", i, values[i])
	}
	// 测试containers
	flag := list.Containers(1, 2)
	assert.True(t, flag == true, "excepting true but found %v", flag)
	// 测试swap
	list.Swap(0, 1)
	v1, _ := list.Get(0)
	assert.True(t, v1 == 1, "excepting 1 but found %v", v1)
	v2, _ := list.Get(1)
	assert.True(t, v2 == 0, "excepting 0 but found %v", v2)
	// 测试insert
	list.Insert(0, -1)
	var res = []int{-1, 1, 0, 2, 3}
	values = list.Values()
	for i := 0; i < len(res); i++ {
		assert.True(t, values[i] == res[i], "excepting %v but found %v", res[i], values[i])
	}
	// 测试set
	list.Set(0, 1)
	values = list.Values()
	assert.True(t, values[0] == 1, "excepting 1 but found %v", values[0])

}
