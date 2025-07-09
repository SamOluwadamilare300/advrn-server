package models

import (
	"time"
	"gorm.io/gorm"
)

type VerificationType string

const (
	VerificationIdentity   VerificationType = "identity"
	VerificationEmployment VerificationType = "employment"
	VerificationLandlord   VerificationType = "landlord"
	VerificationEmployer   VerificationType = "employer"
)

type VerificationStatus string

const (
	VerificationPending  VerificationStatus = "pending"
	VerificationApproved VerificationStatus = "approved"
	VerificationRejected VerificationStatus = "rejected"
	VerificationExpired  VerificationStatus = "expired"
)

type UserVerification struct {
	ID               uint               `json:"id" gorm:"primaryKey"`
	UserID           uint               `json:"user_id" gorm:"not null"`
	Type             VerificationType   `json:"type" gorm:"not null"`
	Status           VerificationStatus `json:"status" gorm:"default:'pending'"`
	SubmittedData    string             `json:"submitted_data" gorm:"type:text"` // JSON data
	ReviewNotes      string             `json:"review_notes" gorm:"type:text"`
	ReviewedBy       *uint              `json:"reviewed_by"`
	ExpiresAt        *time.Time         `json:"expires_at"`
	SubmittedAt      time.Time          `json:"submitted_at"`
	ReviewedAt       *time.Time         `json:"reviewed_at"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
	DeletedAt        gorm.DeletedAt     `json:"-" gorm:"index"`

	// Relationships
	User       User                      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Reviewer   *User                     `json:"reviewer,omitempty" gorm:"foreignKey:ReviewedBy"`
	Documents  []VerificationDocument    `json:"documents,omitempty" gorm:"foreignKey:VerificationID"`
}

type VerificationDocument struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	VerificationID uint   `json:"verification_id" gorm:"not null"`
	DocumentType   string `json:"document_type" validate:"required"`
	FileName       string `json:"file_name" validate:"required"`
	FileURL        string `json:"file_url" validate:"required"`
	FileSize       int64  `json:"file_size"`
	MimeType       string `json:"mime_type"`
	UploadedAt     time.Time `json:"uploaded_at"`
	CreatedAt      time.Time `json:"created_at"`

	// Relationships
	Verification UserVerification `json:"verification,omitempty" gorm:"foreignKey:VerificationID"`
}
