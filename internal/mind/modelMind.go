package mind

import (
	"fmt"
	"greeschenko/workoverlord2/internal/interfaces"
	"greeschenko/workoverlord2/internal/models"
)

type MIND struct {
	Cells   map[string]*models.Cell     `json:"cells"`
	Storage interfaces.StorageInterface `json:"-"`
}

func NewMIND(storageints interfaces.StorageInterface) *MIND {
	return &MIND{
		Cells:   make(map[string]*models.Cell),
		Storage: storageints,
	}
}

func (m *MIND) GetAll() map[string]*models.Cell {
	return m.Cells
}

func (m *MIND) GetOne(key string) (*models.Cell, error) {
	if c, exist := m.Cells[key]; exist {
		return c, nil
	}
	return nil, fmt.Errorf("failed to find cell with key: %v", key)
}

func (m *MIND) Add(key string, cell models.Cell) (*models.Cell, error) {
	m.Cells[key] = &cell
	return m.Cells[key], nil
}

func (m *MIND) Patch(key string, cell models.Cell) (*models.Cell, error) {
	m.Cells[key] = &cell
	return m.Cells[key], nil
}

func (m *MIND) Delete(key string) error {
	delete(m.Cells, key)
	return nil
}
