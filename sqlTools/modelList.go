package sqlTools

import (
	"time"
)

type IrsCmsSync struct {
	ID         uint   `gorm:"primaryKey"`
	PrimaryID  string `gorm:"uniqueIndex;not null"`
	Type       string `gorm:"uniqueIndex;not null"`
	Utime      uint
	Content    string
	Status     string
	CreateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Reason     string
	Action     string
}

func (i *IrsCmsSync) TableName() string {
	return "irs_cms_sync"
}
