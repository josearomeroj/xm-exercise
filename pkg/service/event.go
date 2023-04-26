package service

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Action uint32

const (
	EventCreate Action = iota
	EventUpdate
	EventRemove
)

type event struct {
	PerformerId uuid.UUID `json:"performer_id"`
	TargetId    uuid.UUID `json:"target_id"`
	EventType   Action    `json:"event_type"`
}

func (e *event) bytes() []byte {
	b, _ := json.Marshal(e)
	return b
}
