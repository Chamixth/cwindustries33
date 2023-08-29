package dto

type Projects struct {
    ObjectId string `gorm:"column:objectId" json:"objectId" validate:"required"`
    Name string `json:"name" validate:"required"`
    Description string `json:"description" validate:"required"`
    StartDate string `json:"startDate" validate:"required"`
    EndDate string `json:"endDate" validate:"required"`
    Status string `json:"status" validate:"required"`
    Client string `json:"client" validate:"required"`
    Budget float64 `json:"budget" validate:"required"`
    TeamMembers  `json:"teamMembers" validate:"required"`
    Attachments  `json:"attachments" validate:"required"`
    }

