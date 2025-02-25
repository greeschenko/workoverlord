package interfaces

import (
	"greeschenko/workoverlord2/internal/models"
)

type GUIInterface interface {
	Start()
}

type StorageInterface interface {
	SetSecret(string)
	Load() error
	Save()
}

type DataInterface interface {
	GetAll() map[string]*models.Cell
	GetOne(string) (*models.Cell, error)
	Add(string, models.Cell) (*models.Cell, error)
	Patch(string, models.Cell) (*models.Cell, error)
	Delete(string) error
}
