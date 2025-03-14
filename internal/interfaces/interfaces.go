package interfaces

import (
	"greeschenko/workoverlord2/internal/models"
)

type GUIInterface interface {
	Start()
}

type StorageInterface interface {
	SetSecret(string)
	Load() ([]byte, error)
	Save([]byte)
}

type DataInterface interface {
	SetSecret(string) error
	Load() error
	GetAll() map[string]*models.Cell
	GetOne(string) (*models.Cell, error)
	Add(string, models.Cell) (*models.Cell, error)
	Patch(string, models.Cell) (*models.Cell, error)
	Delete(string) error
}
