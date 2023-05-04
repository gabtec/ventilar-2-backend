package types

import "time"

// Ward is a representation of wards table ERROR.
type Ward struct {
	ID        uint   `json:"ward_id" gorm:"primaryKey"`
	Name      string `json:"name" validate:"required"`
	BelongsTo string `json:"belongs_to"`

	IsPark bool `json:"is_park" gorm:"default:false"`

	// // Relations
	// Users 		[]User	`json:"users" gorm:"constraint:OnDelete:SET NULL;"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
