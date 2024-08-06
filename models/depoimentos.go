package models

import "gorm.io/gorm"

type Depoimentos struct {
	gorm.Model
	Autor string `gorm:"size:255;not null" json:"autor"`
	Texto string `gorm:"size:255;not null" json:"texto"`
	Foto  string `gormL:"size:255;not null" json:"foto"` // Sera um link para a foto
}
