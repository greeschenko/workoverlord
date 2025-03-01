package mind

import (
	"encoding/json"
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

func (m *MIND) SetSecret(secret string) error {
	m.Storage.SetSecret(secret)
	return nil
}

func (m *MIND) Load() error {
	data, err := m.Storage.Load()
	if err != nil {
		return err
	}
	json.Unmarshal(data, &m)

	return nil
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
	m.saveData()
	return m.Cells[key], nil
}

func (m *MIND) Patch(key string, updates models.Cell) (*models.Cell, error) {
	existingCell, exists := m.Cells[key]
	if !exists {
		return nil, fmt.Errorf("cell with key %s not found", key)
	}

	if updates.Content != "" {
		existingCell.Content = updates.Content
	}
	if updates.Position != nil {
		existingCell.Position = updates.Position
	}
	if updates.Size != nil {
		existingCell.Size = updates.Size
	}
	if updates.Status != nil {
		existingCell.Status = updates.Status
	}
	if updates.Style != nil {
		existingCell.Style = updates.Style
	}

	m.saveData()

	return existingCell, nil
}

func (m *MIND) Delete(key string) error {
	delete(m.Cells, key)
	m.saveData()
	return nil
}

func (m *MIND) saveData() {
	USERMINDjson, err := json.MarshalIndent(m, " ", " ")
	if err != nil {
		panic(err)
	}
	m.Storage.Save(USERMINDjson)
}
