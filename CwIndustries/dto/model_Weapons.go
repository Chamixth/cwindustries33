package dto

type Weapons struct {
    ObjectId string `gorm:"column:objectId" json:"objectId" validate:"required"`
    Name string `json:"name" validate:"required"`
    Type string `json:"type" validate:"required"`
    Damage int `json:"damage" validate:"required"`
    Range string `json:"range" validate:"required"`
    Price float64 `json:"price" validate:"required"`
    Ammo int `json:"ammo" validate:"required"`
    IsAvailable bool `json:"isAvailable" validate:"required"`
    }

