package dto

type Engineers struct {
    ObjectId string `gorm:"column:objectId" json:"objectId" validate:"required"`
    Name string `json:"name" validate:"required"`
    Age int `json:"age" validate:"required"`
    Email string `json:"email" validate:"required"`
    Phone string `json:"phone" validate:"required"`
    Address string `json:"address" validate:"required"`
    Specialization string `json:"specialization" validate:"required"`
    Experience int `json:"experience" validate:"required"`
    Salary float64 `json:"salary" validate:"required"`
    Gender string `json:"gender" validate:"required"`
    }

