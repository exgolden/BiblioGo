package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Nombre           string `gorm:"not null" json:"nombre"`
	Apellido_Paterno string `json:"apellido_paterno"`
	Apellido_Materno string `json:"apellido_materno"`
	Boleta           string `gorm:"not null;unique_index" json:"boleta"`
	Email            string `gorm:"not null" json:"email"`
	Plantel          string `gorm:"not null" json:"plantel"`
	Prestamos        uint   `gorm:"type:smallint;default:0" json:"prestamos"`
	Activo           bool   `gorm:"default:false" json:"activo"`
	Bloqueado        bool   `gorm:"default:false" json:"bloqueado"`
	Multa            uint   `gorm:"type:smallint;default:0" json:"multa"`
	SuperUser        bool   `gorm:"default:false" json:"superuser"`
}
