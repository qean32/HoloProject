package model

type EventFunction func(event Event)
type FnRerutnEvent func(arr []string) (event Event, _error bool)
