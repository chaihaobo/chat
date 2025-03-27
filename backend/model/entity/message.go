package entity

import "encoding/json"

type Message struct {
	BaseEntity
	From    uint64 `gorm:"column:from"`
	To      uint64 `gorm:"column:to"`
	Content string `gorm:"column:content"`
}

func (s Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Message) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

func (Message) TableName() string {
	return "message"
}
