package dto

type Scientists struct {
    ObjectId string `gorm:"column:objectId" json:"objectId" validate:"required"`
    Name string `json:"name" validate:"required"`
    Age int `json:"age" validate:"required"`
    Gender string `json:"gender" validate:"required"`
    BirthDate string `json:"birthDate" validate:"required"`
    Email string `json:"email" validate:"required"`
    Specialization string `json:"specialization" validate:"required"`
    Country string `json:"country" validate:"required"`
    MobileNumber string `json:"mobileNumber" validate:"required"`
    Designation string `json:"designation" validate:"required"`
    }

