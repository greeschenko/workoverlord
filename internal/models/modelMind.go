package models

type MIND struct {
	Cells map[string]*Cell `json:"cells"`
}

func NewMIND() *MIND {
	return &MIND{
		Cells: make(map[string]*Cell),
	}
}
