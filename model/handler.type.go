package model

type Handler func(event Event) error
type ParseEvent func(arr []string) (event Event, err error)
