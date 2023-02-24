package crevendo

import "github.com/crevendo/crevendo/data"

func GetFieldValueByName(name string, fields []*data.Field) string {
	for i := range fields {
		if fields[i].Name == name {
			return fields[i].Value
		}
	}
	return ""
}
