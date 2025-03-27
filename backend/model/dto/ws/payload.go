package ws

import "github.com/mitchellh/mapstructure"

type (
	PayloadData map[string]any
	Payload     struct {
		Event EventType   `json:"event"`
		Data  PayloadData `json:"data"`
	}
)

func NewPayload(event EventType, data any) *Payload {
	payloadData := make(map[string]any)
	mapstructure.Decode(data, &payloadData)
	return &Payload{
		Event: event,
		Data:  payloadData,
	}
}

func (p Payload) ScanDataTo(target any) {
	mapstructure.Decode(p.Data, target)
}
