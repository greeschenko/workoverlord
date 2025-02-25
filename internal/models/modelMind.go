package models

import (
	"fmt"
)

type MIND struct {
	Cells map[string]*Cell `json:"cells"`
}

func NewMIND() *MIND {
	return &MIND{
		Cells: make(map[string]*Cell),
	}
}

func (m *MIND) GetAll() map[string]*Cell {
	return m.Cells
}

func (m *MIND) GetOne(key string) (*Cell, error) {
	if c, exist := m.Cells[key]; exist {
		return c, nil
	}
	return nil, fmt.Errorf("failed to find cell with key: %v", key)
}

func (m *MIND) Add(key string, cell Cell) (*Cell, error) {
	m.Cells[key] = &cell
	return m.Cells[key], nil
}

func (m *MIND) Patch(key string, cell Cell) (*Cell, error) {
	m.Cells[key] = &cell
	return m.Cells[key], nil
}

func (m *MIND) Delete(key string) error {
	delete(m.Cells, key)
	return nil
}
