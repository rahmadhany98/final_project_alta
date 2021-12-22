package model

import "gorm.io/datatypes"

type Transaction struct {
	ID          int            `json:"id" form:"id"`
	Date        datatypes.Date `json:"date" form:"date"`
	Amount      int            `json:"amount" form:"amount"`
	Description string         `json:"description" form:"description"`
	Category    string         `json:"category" form:"category"`
	UserID      int            `json:"userid" form:"userid"`
	User        User           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
