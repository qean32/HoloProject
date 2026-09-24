package model

type FieldType struct {
	Message string

	Position      int
	PositionRange int

	Prefix    string
	PrefixSRG []int
}

type FieldPayload struct {
	Prefix    string
	PrefixSRG []int
}
