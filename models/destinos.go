package models

import "gorm.io/gorm"

type Destinos struct {
	gorm.Model
	Foto            string  `gorm:"size:255;not null" json:"foto"`
	Foto2           string  `gorm:"size:255;" json:"foto2"`
	Nome            string  `gorm:"size:255;not null" json:"nome"`
	Preco           float64 `gorm:"not null" json:"preco"`
	Meta            string  `gorm:"size:160;" json:"meta"`
	TextoDescritivo string  `gorm:"size:255" json:"texto_descritivo"`
}
