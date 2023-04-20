package data

import (
	"strconv"

	"github.com/codebdy/entify/shared"
)

func ConvertId(id interface{}) uint64 {
	switch id.(type) {
	case string:
		id, err := strconv.ParseUint(id.(string), 10, 64)
		if err != nil {
			panic("Convert id error:" + err.Error())
		}
		return id
	}

	return id.(uint64)
}

func ConvertObjectId(object map[string]interface{}) map[string]interface{} {
	if object[shared.ID] == nil {
		return object
	}
	switch object[shared.ID].(type) {
	case string:
		id, err := strconv.ParseUint(object[shared.ID].(string), 10, 64)
		if err != nil {
			panic("Convert id error:" + err.Error())
		}

		object[shared.ID] = id
	}

	return object
}
