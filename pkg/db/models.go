package db

import (
	"time"
)

type TransmitterData struct {
	ID        uint      `gorm:"primaryKey"`
	SessionID string    `gorm:"column:session_id"`
	Frequency float64   `gorm:"column:frequency"`
	Timestamp time.Time `gorm:"column:timestamp"`
}
