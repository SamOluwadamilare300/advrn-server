package models

import (
	"time"
	"gorm.io/gorm"
)

type VirtualTourType string

const (
	TourType360Photos VirtualTourType = "360_photos"
	TourTypeVideo     VirtualTourType = "video"
	TourTypeLive      VirtualTourType = "live"
)

type VirtualTour struct {
	ID           uint            `json:"id" gorm:"primaryKey"`
	PropertyID   uint            `json:"property_id" gorm:"not null"`
	Type         VirtualTourType `json:"type" gorm:"not null"`
	Title        string          `json:"title" validate:"required"`
	Description  string          `json:"description" gorm:"type:text"`
	IsActive     bool            `json:"is_active" gorm:"default:true"`
	ViewCount    int             `json:"view_count" gorm:"default:0"`
	Duration     int             `json:"duration"` 
	TourData     string          `json:"tour_data" gorm:"type:text"` 
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeletedAt    gorm.DeletedAt  `json:"-" gorm:"index"`

	// Relationships
	Property     Property        `json:"property,omitempty" gorm:"foreignKey:PropertyID"`
	Images360    []Tour360Image  `json:"images_360,omitempty" gorm:"foreignKey:VirtualTourID"`
	Videos       []TourVideo     `json:"videos,omitempty" gorm:"foreignKey:VirtualTourID"`
	LiveSessions []LiveTourSession `json:"live_sessions,omitempty" gorm:"foreignKey:VirtualTourID"`
}

type Tour360Image struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	VirtualTourID uint      `json:"virtual_tour_id" gorm:"not null"`
	RoomName      string    `json:"room_name" validate:"required"`
	ImageURL      string    `json:"image_url" validate:"required"`
	ThumbnailURL  string    `json:"thumbnail_url"`
	FileSize      int64     `json:"file_size"`
	Width         int       `json:"width"`
	Height        int       `json:"height"`
	Order         int       `json:"order" gorm:"default:0"`
	Hotspots      string    `json:"hotspots" gorm:"type:text"` // JSON array of hotspot data
	CreatedAt     time.Time `json:"created_at"`

	// Relationships
	VirtualTour VirtualTour `json:"virtual_tour,omitempty" gorm:"foreignKey:VirtualTourID"`
}

type TourVideo struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	VirtualTourID uint      `json:"virtual_tour_id" gorm:"not null"`
	VideoURL      string    `json:"video_url" validate:"required"`
	ThumbnailURL  string    `json:"thumbnail_url"`
	Duration      int       `json:"duration"` // in seconds
	FileSize      int64     `json:"file_size"`
	Quality       string    `json:"quality" gorm:"default:'720p'"`
	CreatedAt     time.Time `json:"created_at"`

	// Relationships
	VirtualTour VirtualTour `json:"virtual_tour,omitempty" gorm:"foreignKey:VirtualTourID"`
}

type LiveTourSession struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	VirtualTourID uint      `json:"virtual_tour_id" gorm:"not null"`
	HostID        uint      `json:"host_id" gorm:"not null"` // Landlord/Agent
	ScheduledAt   time.Time `json:"scheduled_at" validate:"required"`
	Duration      int       `json:"duration" gorm:"default:30"` // in minutes
	MeetingURL    string    `json:"meeting_url"`
	MeetingID     string    `json:"meeting_id"`
	Status        string    `json:"status" gorm:"default:'scheduled'"` // scheduled, active, completed, cancelled
	MaxAttendees  int       `json:"max_attendees" gorm:"default:10"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	// Relationships
	VirtualTour VirtualTour `json:"virtual_tour,omitempty" gorm:"foreignKey:VirtualTourID"`
	Host        User        `json:"host,omitempty" gorm:"foreignKey:HostID"`
	Attendees   []LiveTourAttendee `json:"attendees,omitempty" gorm:"foreignKey:LiveTourSessionID"`
}

type LiveTourAttendee struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	LiveTourSessionID uint      `json:"live_tour_session_id" gorm:"not null"`
	UserID            uint      `json:"user_id" gorm:"not null"`
	JoinedAt          *time.Time `json:"joined_at"`
	LeftAt            *time.Time `json:"left_at"`
	Status            string    `json:"status" gorm:"default:'registered'"` // registered, joined, left
	CreatedAt         time.Time `json:"created_at"`

	// Relationships
	LiveTourSession LiveTourSession `json:"live_tour_session,omitempty" gorm:"foreignKey:LiveTourSessionID"`
	User            User            `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
