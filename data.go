package crevendo

import "github.com/crevendo/crevendo/data"

func GetDataValue(key string, data []*data.Data) string {
	var value string
	for i := range data {
		if key == data[i].Key {
			value = data[i].Value
		}
	}
	return value
}
