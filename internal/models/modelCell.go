package models

type CellStatus int
type CellType int

// Define constants for the possible statuses
const (
	CellStatusConfig   CellStatus = iota // 0
	CellStatusActive                     // 1
	CellStatusArchived                   // 2
	CellStatusDeleted                    // 3
)

// Define constants for the possible statuses
const (
	CellTypeText CellType = iota // 0
	CellTypeImg                  // 1
)

type Cell struct {
	Content  string      `json:"content"`
	Position *[2]int     `json:"position"`
	Size     *[2]int     `json:"size"`
	Status   *CellStatus `json:"status"`
	Type     *CellType   `json:"type"`
	Style    *Style      `json:"style"`
	//Synapses map[string]*Synapse `json:"synaptises"`
}
