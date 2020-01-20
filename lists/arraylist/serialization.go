package arraylist

import "encoding/json"

/**
 * @Author: eggsy
 * @Description:
 * @File:  serialization
 * @Version: 1.0.0
 * @Date: 2020-01-20 14:17
 */

func (list *List) ToJSON() ([]byte, error) {
	return json.Marshal(list.elements[:list.size])
}

func (list *List) FromJSON(data []byte) error {
	err := json.Unmarshal(data, &list.elements)
	if err == nil {
		list.size = len(list.elements)
	}
	return err
}
