package data

import (
	"strconv"

	"github.com/codebdy/entify-core/shared"
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
	if object[shared.ID_NAME] == nil {
		return object
	}
	switch object[shared.ID_NAME].(type) {
	case string:
		id, err := strconv.ParseUint(object[shared.ID_NAME].(string), 10, 64)
		if err != nil {
			panic("Convert id error:" + err.Error())
		}

		object[shared.ID_NAME] = id
	}

	return object
}
