package utils

import "time"

/**
 * @Author: eggsy
 * @Description: 实现两个接口到比较 a<b 负数,a=b 0,a>b 正数
 * @File:  comparator
 * @Version: 1.0.0
 * @Date: 2019-10-22 10:27
 */

type Comparator func(a, b interface{}) int

// 提供字符串类型的比较
// 将字符串转换成byte数组，按位进行比较，
// 若a[i]>b[i]，则a>b
// 若a[i]<b[i] 则a<b
// 若现有位数a[i]=b[i] 则比较长度，谁长谁大
//s1 := a.(string)
//s2 := b.(string)
//min := len(s1)
//if min > len(s2) {
//	min = len(s2)
//}
////按位进行比较
//for i := 0; i < min; i++ {
//	switch {
//	case int(s1[i])-int(s2[i]) > 0:
//		return 1
//	case int(s1[i])-int(s2[i]) < 0:
//		return -1
//	}
//}
//return len(s1) - len(s2)
func StringComparator(a, b interface{}) int {
	s1 := a.(string)
	s2 := b.(string)
	switch {
	case s1 > s2:
		return 1
	case s1 < s2:
		return -1
	default:
		return 0
	}
}

func IntComparator(a, b interface{}) int {
	aAsserted := a.(int)
	bAsserted := b.(int)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func Int8Comparator(a, b interface{}) int {
	aAsserted := a.(int8)
	bAsserted := b.(int8)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func Int16Comparator(a, b interface{}) int {
	aAsserted := a.(int16)
	bAsserted := b.(int16)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func Int32Comparator(a, b interface{}) int {
	aAsserted := a.(int32)
	bAsserted := b.(int32)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func Int64Comparator(a, b interface{}) int {
	aAsserted := a.(int64)
	bAsserted := b.(int64)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func UintComparator(a, b interface{}) int {
	aAsserted := a.(uint)
	bAsserted := b.(uint)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func Uint8Comparator(a, b interface{}) int {
	aAsserted := a.(uint8)
	bAsserted := b.(uint8)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
func Uint16Comparator(a, b interface{}) int {
	aAsserted := a.(uint16)
	bAsserted := b.(uint16)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func Uint32Comparator(a, b interface{}) int {
	aAsserted := a.(uint32)
	bAsserted := b.(uint32)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func Uint64Comparator(a, b interface{}) int {
	aAsserted := a.(uint64)
	bAsserted := b.(uint64)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
func Float64Comparator(a, b interface{}) int {
	aAsserted := a.(float64)
	bAsserted := b.(float64)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func Float32Comparator(a, b interface{}) int {
	aAsserted := a.(float32)
	bAsserted := b.(float32)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func ByteComparator(a, b interface{}) int {
	aAsserted := a.(byte)
	bAsserted := b.(byte)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func RuneComparator(a, b interface{}) int {
	aAsserted := a.(rune)
	bAsserted := b.(rune)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func TimeComparator(a, b interface{}) int {
	aAsserted := a.(time.Time)
	bAsserted := b.(time.Time)
	switch {
	case aAsserted.After(bAsserted):
		return 1
	case aAsserted.Before(bAsserted):
		return -1
	default:
		return 0
	}
}
