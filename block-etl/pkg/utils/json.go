package utils

import "encoding/json"

func PrettyStruct(data interface{}) string {
	val, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err.Error()
	}
	return string(val)
}
