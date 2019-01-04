package constants

import "time"

type Task struct {
	ID         string    `json:"id"`
	Channel    string    `json:"channel"`
	Subject    string    `json:"subject"`
	Data       []byte    `json:"data,omitempty"`
	CreateTime time.Time `json:"createTime"`
}
