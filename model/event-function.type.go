package model

type EventFunction func(e Event)
type FnReturnEvent func(arr []string) (e Event, _error bool)
