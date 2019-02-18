package api

import (
	// "github.com/jinzhu/gorm"
	"time"
)

type Book struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time `json:"-"`
	Name      string    `gorm:"size:255;unique_index" json:"name"`
	Contacts  []Contact `json:"contacts,omitempty"` // One-To-Many relationship
	Active    bool      `json:"-"`
}

type Contact struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time `json:"-"`
	BookID    uint      `gorm:"type:int REFERENCES books(id)" json:"book_id,omitempty"`
	Name      string    `gorm:"size:255" json:"name;omitempty"`
	Email     string    `gorm:"type:varchar(100);unique_index" json:"email_id,omitempty"`
	Active    bool      `json:"-"`
}
