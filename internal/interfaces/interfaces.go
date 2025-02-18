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
	GetSecret() [32]byte
	Load() error
	Save()
	GetData() models.MIND
}
