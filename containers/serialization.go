package containers

/**
 * @Author: eggsy
 * @Description:
 * @File:  serialization
 * @Version: 1.0.0
 * @Date: 2020-01-06 19:21
 */

type JSONSerializer interface {
	ToJSON() ([]byte, error)
}

type JSONDeserializer interface {
	FromJSON([]byte) error
}
