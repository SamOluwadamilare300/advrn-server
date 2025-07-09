package models

import (
	"time"
	"gorm.io/gorm"
)

type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "pending"
	PaymentCompleted PaymentStatus = "completed"
	PaymentFailed    PaymentStatus = "failed"
	PaymentRefunded  PaymentStatus = "refunded"
)

type PaymentType string

const (
	PaymentRent     PaymentType = "rent"
	PaymentDeposit  PaymentType = "deposit"
	PaymentFee      PaymentType = "fee"
)

type Payment struct {
	ID                uint          `json:"id" gorm:"primaryKey"`
	UserID            uint          `json:"user_id" gorm:"not null"`
	PropertyID        uint          `json:"property_id" gorm:"not null"`
	Amount            float64       `json:"amount" gorm:"not null"`
	Currency          string        `json:"currency" gorm:"default:'NGN'"`
	PaymentType       PaymentType   `json:"payment_type" gorm:"not null"`
	Status            PaymentStatus `json:"status" gorm:"default:'pending'"`
	PaymentReference  string        `json:"payment_reference" gorm:"unique"`
	ProviderReference string        `json:"provider_reference"`
	Provider          string        `json:"provider" gorm:"default:'paystack'"`
	Description       string        `json:"description"`
	Metadata          string        `json:"metadata" gorm:"type:text"`
	PaidAt            *time.Time    `json:"paid_at"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`

	// Relationships
	User     User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Property Property `json:"property,omitempty" gorm:"foreignKey:PropertyID"`
}

type RecurringPayment struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	UserID           uint      `json:"user_id" gorm:"not null"`
	PropertyID       uint      `json:"property_id" gorm:"not null"`
	Amount           float64   `json:"amount" gorm:"not null"`
	Currency         string    `json:"currency" gorm:"default:'NGN'"`
	Frequency        string    `json:"frequency" gorm:"default:'monthly'"` // monthly, quarterly, yearly
	NextPaymentDate  time.Time `json:"next_payment_date"`
	IsActive         bool      `json:"is_active" gorm:"default:true"`
	AuthorizationCode string   `json:"authorization_code"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	// Relationships
	User     User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Property Property `json:"property,omitempty" gorm:"foreignKey:PropertyID"`
}

type EscrowAccount struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	PropertyID uint    `json:"property_id" gorm:"not null"`
	TenantID   uint    `json:"tenant_id" gorm:"not null"`
	LandlordID uint    `json:"landlord_id" gorm:"not null"`
	Amount     float64 `json:"amount" gorm:"not null"`
	Status     string  `json:"status" gorm:"default:'held'"` // held, released, disputed
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	// Relationships
	Property Property `json:"property,omitempty" gorm:"foreignKey:PropertyID"`
	Tenant   User     `json:"tenant,omitempty" gorm:"foreignKey:TenantID"`
	Landlord User     `json:"landlord,omitempty" gorm:"foreignKey:LandlordID"`
}
