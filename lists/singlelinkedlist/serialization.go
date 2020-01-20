package linkedlist

import "encoding/json"

/**
 * @Author: eggsy
 * @Description:
 * @File:  serialization
 * @Version: 1.0.0
 * @Date: 2020-01-20 16:42
 */
func (list *List) ToJSON() ([]byte, error) {
	return json.Marshal(list.Values())
}

func (list *List) FromJSON(data []byte) error {
	elements := []interface{}{}
	err := json.Unmarshal(data, &elements)
	if err == nil {
		list.Clear()
		list.Add(elements...)
	}
	return err
}
