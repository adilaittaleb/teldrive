package models

import (
	"time"

	"github.com/divyam234/teldrive/pkg/schemas"
	"gorm.io/datatypes"
)

type File struct {
	Id        string                            `gorm:"type:uuid;primaryKey;default:uuid7()"`
	Name      string                            `gorm:"type:text;not null"`
	Type      string                            `gorm:"type:text;not null"`
	MimeType  string                            `gorm:"type:text;not null"`
	Size      *int64                            `gorm:"type:bigint"`
	Starred   bool                              `gorm:"default:false"`
	Category  string                            `gorm:"type:text"`
	Encrypted bool                              `gorm:"default:false"`
	UserID    int64                             `gorm:"type:bigint;not null"`
	Status    string                            `gorm:"type:text"`
	ParentID  string                            `gorm:"type:uuid;index"`
	Parts     datatypes.JSONSlice[schemas.Part] `gorm:"type:jsonb"`
	ChannelID *int64                            `gorm:"type:bigint"`
	CreatedAt time.Time                         `gorm:"default:timezone('utc'::text, now())"`
	UpdatedAt time.Time                         `gorm:"default:timezone('utc'::text, now())"`
}
