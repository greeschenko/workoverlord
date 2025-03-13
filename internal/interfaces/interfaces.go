package interfaces

import (
	"greeschenko/workoverlord2/internal/models"
)

type GUIInterface interface {
	Start()
}

type SpatialObject interface {
	ID() string
	Coordinates() [2]int
}

type Positioner interface {
	NearestNeighbor(target [2]int) string
	FindNearestInDirection(target SpatialObject, direction string) string
	Rebuild(objects []SpatialObject)
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
