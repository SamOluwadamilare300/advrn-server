package models

import (
	"time"
	"gorm.io/gorm"
)

type ApplicationStatus string

const (
	ApplicationSubmitted   ApplicationStatus = "submitted"
	ApplicationUnderReview ApplicationStatus = "under_review"
	ApplicationApproved    ApplicationStatus = "approved"
	ApplicationRejected    ApplicationStatus = "rejected"
	ApplicationWithdrawn   ApplicationStatus = "withdrawn"
)

type RentalApplication struct {
	ID                uint              `json:"id" gorm:"primaryKey"`
	UserID            uint              `json:"user_id" gorm:"not null"`
	PropertyID        uint              `json:"property_id" gorm:"not null"`
	Status            ApplicationStatus `json:"status" gorm:"default:'submitted'"`
	
	// Personal Information
	FullName          string    `json:"full_name" validate:"required"`
	Email             string    `json:"email" validate:"required,email"`
	Phone             string    `json:"phone" validate:"required"`
	DateOfBirth       time.Time `json:"date_of_birth" validate:"required"`
	
	// Employment Information
	EmploymentStatus  string  `json:"employment_status" validate:"required"`
	EmployerName      string  `json:"employer_name"`
	JobTitle          string  `json:"job_title"`
	MonthlyIncome     float64 `json:"monthly_income" validate:"required,min=0"`
	EmploymentLength  string  `json:"employment_length"`
	
	// Previous Address
	PreviousAddress   string `json:"previous_address"`
	PreviousLandlord  string `json:"previous_landlord"`
	ReasonForMoving   string `json:"reason_for_moving"`
	
	// References
	References        string `json:"references" gorm:"type:text"` // JSON array
	
	// Additional Information
	PetInformation    string `json:"pet_information"`
	AdditionalNotes   string `json:"additional_notes" gorm:"type:text"`
	
	// Application Score
	Score             float64 `json:"score" gorm:"default:0"`
	
	// Timestamps
	SubmittedAt       time.Time `json:"submitted_at"`
	ReviewedAt        *time.Time `json:"reviewed_at"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	User     User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Property Property `json:"property,omitempty" gorm:"foreignKey:PropertyID"`
	Documents []ApplicationDocument `json:"documents,omitempty" gorm:"foreignKey:ApplicationID"`
}

type ApplicationDocument struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	ApplicationID uint   `json:"application_id" gorm:"not null"`
	DocumentType  string `json:"document_type" validate:"required"` // id_card, payslip, bank_statement, etc.
	FileName      string `json:"file_name" validate:"required"`
	FileURL       string `json:"file_url" validate:"required"`
	FileSize      int64  `json:"file_size"`
	MimeType      string `json:"mime_type"`
	UploadedAt    time.Time `json:"uploaded_at"`
	CreatedAt     time.Time `json:"created_at"`

	// Relationships
	Application RentalApplication `json:"application,omitempty" gorm:"foreignKey:ApplicationID"`
}
