package stacks

import "github.com/Eggsyz/pangu/containers"

/**
 * @Author: eggsy
 * @Description: golang 实现栈
 * @File:  stack
 * @Version: 1.0.0
 * @Date: 2019-10-17 17:15
 */
type Stack interface {
	Push(value interface{})
	Pop() (value interface{}, ok bool)
	Peek() (value interface{}, ok bool)
	containers.Container
}
