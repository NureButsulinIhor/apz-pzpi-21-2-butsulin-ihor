package models

import (
	"gorm.io/gorm"
	"time"
)

type Timetable struct {
	gorm.Model
	WorkerID  uint      `json:"worker_id"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
