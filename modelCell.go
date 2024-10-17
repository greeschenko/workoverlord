package main

type CellStatus int

// Define constants for the possible statuses
const (
	CellStatusConfig   CellStatus = iota // 0
	CellStatusActive                     // 1
	CellStatusArchived                   // 2
	CellStatusDeleted                    // 3
)

type Cell struct {
	Content    string              `json:"content"`
	Position   [2]int              `json:"position"`
	Size       [2]int              `json:"size"`
	Status     CellStatus          `json:"status"`
	Style      *Style              `json:"style"`
	Synapses map[string]*Synapse `json:"synaptises"`
}
