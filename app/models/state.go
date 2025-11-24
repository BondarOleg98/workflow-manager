package models

type State string

const (
	CREATED State = "CREATED"
	SUCCESS State = "SUCCESS"
	ERROR   State = "ERROR"
	IDLE    State = "IDLE"
	RUNNING State = "RUNNING"
)
