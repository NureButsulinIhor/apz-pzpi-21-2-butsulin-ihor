package models

import (
	"gorm.io/gorm"
	"time"
)

type WeighingResult struct {
	gorm.Model
	SlotID uint      `json:"slot_id"`
	Weight float64   `json:"weight"`
	Time   time.Time `json:"time"`
}
