package models

import (
	"time"
	"gorm.io/gorm"
)

type LeaseStatus string

const (
	LeaseDraft      LeaseStatus = "draft"
	LeaseActive     LeaseStatus = "active"
	LeaseExpired    LeaseStatus = "expired"
	LeaseTerminated LeaseStatus = "terminated"
	LeaseRenewed    LeaseStatus = "renewed"
)

type Lease struct {
	ID                uint        `json:"id" gorm:"primaryKey"`
	PropertyID        uint        `json:"property_id" gorm:"not null"`
	TenantID          uint        `json:"tenant_id" gorm:"not null"`
	LandlordID        uint        `json:"landlord_id" gorm:"not null"`
	Status            LeaseStatus `json:"status" gorm:"default:'draft'"`
	
	// Lease Terms
	MonthlyRent       float64   `json:"monthly_rent" validate:"required,min=0"`
	SecurityDeposit   float64   `json:"security_deposit" validate:"required,min=0"`
	LeaseStartDate    time.Time `json:"lease_start_date" validate:"required"`
	LeaseEndDate      time.Time `json:"lease_end_date" validate:"required"`
	LeaseDuration     int       `json:"lease_duration"` // in months
	
	// Additional Terms
	UtilitiesIncluded bool   `json:"utilities_included" gorm:"default:false"`
	PetsAllowed       bool   `json:"pets_allowed" gorm:"default:false"`
	SmokingAllowed    bool   `json:"smoking_allowed" gorm:"default:false"`
	SpecialClauses    string `json:"special_clauses" gorm:"type:text"`
	
	// Signatures
	TenantSigned      bool       `json:"tenant_signed" gorm:"default:false"`
	LandlordSigned    bool       `json:"landlord_signed" gorm:"default:false"`
	TenantSignedAt    *time.Time `json:"tenant_signed_at"`
	LandlordSignedAt  *time.Time `json:"landlord_signed_at"`
	
	// Document
	LeaseDocumentURL  string `json:"lease_document_url"`
	
	// Timestamps
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	Property Property `json:"property,omitempty" gorm:"foreignKey:PropertyID"`
	Tenant   User     `json:"tenant,omitempty" gorm:"foreignKey:TenantID"`
	Landlord User     `json:"landlord,omitempty" gorm:"foreignKey:LandlordID"`
	MaintenanceRequests []MaintenanceRequest `json:"maintenance_requests,omitempty" gorm:"foreignKey:LeaseID"`
}

type MaintenanceRequest struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	LeaseID     uint   `json:"lease_id" gorm:"not null"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" gorm:"type:text"`
	Priority    string `json:"priority" gorm:"default:'medium'"` // low, medium, high, urgent
	Status      string `json:"status" gorm:"default:'submitted'"` // submitted, in_progress, completed, cancelled
	Images      string `json:"images" gorm:"type:text"` // JSON array of image URLs
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relationships
	Lease Lease `json:"lease,omitempty" gorm:"foreignKey:LeaseID"`
}
