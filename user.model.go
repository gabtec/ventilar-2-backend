package main

import "time"

// User is a representation of wards table
type User struct {
	ID           uint   `json:"id"      gorm:"primaryKey"`
	Name         string `json:"name"`
	Role         string `json:"role"    default:"consumer"`
	Mec          uint   `json:"mec"     gorm:"unique"`
	PasswordHash string `json:"password_hash"`
	// Relation with t_wards
	WardID uint `json:"ward_id"`
	// This is needed for reverse relation query
	// Ward					Ward		`json:"ward"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
