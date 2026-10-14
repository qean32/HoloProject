package field

import (
	"main/constants"
	"main/model"
)

var field = model.FieldType{
	PrefixSRG: []int{},
	Prefix:    constants.FIELDPREFIX,
}

func set(payload model.FieldPayload) {
	field.Prefix = payload.Prefix
	field.PrefixSRG = payload.PrefixSRG
}
