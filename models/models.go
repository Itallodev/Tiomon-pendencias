package models

import "gorm.io/gorm"

type Pendencia struct {
	gorm.Model
	ID        int    `json:"id"`
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	Status    string `json:"status"`
}

var pendencias []Pendencia
