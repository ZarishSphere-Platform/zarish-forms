package models

// Form represents a dynamic form definition
type Form struct {
	BaseModel

	Name        string `gorm:"size:255;not null;index" json:"name"`
	Title       string `gorm:"size:500" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Version     string `gorm:"size:50" json:"version"`
	Status      string `gorm:"size:50;default:'draft'" json:"status"` // draft, published, archived

	Schema   string `gorm:"type:jsonb;not null" json:"schema"` // JSON schema for the form
	UISchema string `gorm:"type:jsonb" json:"ui_schema"`       // UI rendering hints

	CreatedBy uint `gorm:"index" json:"created_by"`
}

// FormSubmission represents a submitted form response
type FormSubmission struct {
	BaseModel

	FormID      uint   `gorm:"index;not null" json:"form_id"`
	SubmittedBy uint   `gorm:"index;not null" json:"submitted_by"`
	Data        string `gorm:"type:jsonb;not null" json:"data"`           // JSON data matching form schema
	Status      string `gorm:"size:50;default:'submitted'" json:"status"` // submitted, reviewed, approved
}

// TableName overrides
func (Form) TableName() string {
	return "forms"
}

func (FormSubmission) TableName() string {
	return "form_submissions"
}
