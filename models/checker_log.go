package models

import (
	"time"
)

type (
	CheckerLogModel struct {
		ID        int       `json:"id"`
		TopicName string    `json:"topic_name"`
		QueueName string    `json:"queue_name"`
		Message   string    `json:"message"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
