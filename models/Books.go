package models

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Titulo string `gorm:"not null"`
	Autor  string `gorm:"not null"`
	ISBN   string `gorm:"not null;unique_index"`
	Copia  uint   `gorm:"type:smallint"`
	Estado bool   `gorm:"default:false"`
}
