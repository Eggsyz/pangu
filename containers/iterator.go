package containers

/**
 * @Author: eggsy
 * @Description:
 * @File:  iterator
 * @Version: 1.0.0
 * @Date: 2020-01-06 19:43
 */
type Iterator interface {
	// Returns true if the iteration has more elements.
	HasNext() bool
	// Returns the next element in the iteration
	Next() interface{}
	// Removes from the underlying collection the last element returned by this iterator (optional operation).
	Remove()
}
