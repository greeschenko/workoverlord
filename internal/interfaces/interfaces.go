package interfaces

import (
	"greeschenko/workoverlord2/internal/models"
)

// GUIInterface визначає поведінку графічного інтерфейсу
type GUIInterface interface {
	Start()
}

// StorageInterface визначає поведінку сховища даних
type StorageInterface interface {
	SetSecret(string)
	Load() error
	Save()
}

type DataInterface interface {
	GetAll() map[string]*models.Cell
	GetOne(string) (models.Cell, error)
	Add(models.Cell) (models.Cell, error)
	Patch(string, models.Cell) (models.Cell, error)
	Delete(string) (models.Cell, error)
}
