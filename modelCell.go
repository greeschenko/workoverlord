package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"

	"github.com/go-rest-framework/core"
)

type Cell struct {
	ID       string    `json:"id"`
	Data     string    `json:"data"`
	Status   string    `json:"status"`
	Size     [2]int    `json:"size,omitempty"`
	Position [3]int    `json:"position,omitempty"`
	Cells    []Cell    `json:"cells,omitempty"`
	Synapses []Synapse `json:"synapses,omitempty"`
}

type Synapse struct {
	Points   [][3]int `json:"points,omitempty"`
	Size     int      `json:"size,omitempty"`
	Color    string   `json:"color,omitempty"`
	Linetype string   `json:"linetype,omitempty"`
	Endtype  string   `json:"endtype,omitempty"`
}

func (s Cell) getValue(name string) string {
	res := ""
	switch name {
	case "id":
		res = s.ID
	case "data":
		res = s.Data
	case "status":
		res = s.Status
	}
	return res
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (s *Cell) genID(parentid string) {
	resid, _ := randomHex(2)
	if parentid != "0" {
		s.ID = parentid + " " + resid
	} else {
		s.ID = resid
	}
	fmt.Println("GENERATED ID", s.ID)
}

func (s *Cell) RecalculateIds(parentid string) {
    s.genID(parentid)
    for e := range s.Cells {
        s.Cells[e].RecalculateIds(s.ID)
    }
}

func (s *Cell) AppendCell(newdata Cell) {
	(*s).Cells = append((*s).Cells, newdata)
}

func (s *Cell) Update(newdata Cell) {
	(*s) = newdata
}

func (s *Cell) GetCenter() [3]int {
	return [3]int{s.Position[0] + s.Size[0]/2, s.Position[1] + s.Size[1]/2, 0}
}

// return random GRB hex color string
func randomColor() string {
	var r, g, b *big.Int
	r, _ = rand.Int(rand.Reader, big.NewInt(256))
	g, _ = rand.Int(rand.Reader, big.NewInt(256))
	b, _ = rand.Int(rand.Reader, big.NewInt(256))
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

func generateSynapses(first, second Cell) [][3]int {
	points := [][3]int{}
	switch GetRelativePosition(first, second) {
	case "topleft":
		points = [][3]int{
			first.Position,
			{first.Position[0], second.GetCenter()[1], 0},
			{second.Position[0] + second.Size[0], second.Position[1] + second.Size[1]/2, 0},
		}
	case "topcenter":
		points = [][3]int{
			{first.GetCenter()[0], first.Position[1], 0},
			{first.GetCenter()[0], first.Position[1] - (first.Position[1] - second.Position[1] - second.Size[1]) / 2, 0},
			{second.GetCenter()[0], first.Position[1] - (first.Position[1] - second.Position[1] - second.Size[1]) / 2, 0},
			{second.Position[0] + second.Size[0]/2, second.Position[1] + second.Size[1], 0},
		}
	case "topright":
		points = [][3]int{
			{first.Position[0] + first.Size[0], first.Position[1], 0},
			{first.Position[0] + first.Size[0], second.GetCenter()[1], 0},
			{second.Position[0], second.GetCenter()[1], 0},
		}
	case "midleft":
		points = [][3]int{
			{first.Position[0], first.GetCenter()[1], 0},
			{first.Position[0] - (first.Position[0]-second.Position[0]-second.Size[0])/2, first.GetCenter()[1], 0},
			{first.Position[0] - (first.Position[0]-second.Position[0]-second.Size[0])/2, second.GetCenter()[1], 0},
			{second.Position[0] + second.Size[0], second.GetCenter()[1], 0},
		}
	case "midright":
		points = [][3]int{
			{first.Position[0] + first.Size[0], first.GetCenter()[1], 0},
			{second.Position[0] - (second.Position[0]-first.Position[0]-first.Size[0])/2, first.GetCenter()[1], 0},
			{second.Position[0] - (second.Position[0]-first.Position[0]-first.Size[0])/2, second.GetCenter()[1], 0},
			{second.Position[0], second.GetCenter()[1], 0},
		}
	case "bottomleft":
		points = [][3]int{
			{first.Position[0], first.Position[1] + first.Size[1], 0},
			{first.Position[0], second.GetCenter()[1], 0},
			{second.Position[0] + second.Size[0], second.GetCenter()[1], 0},
		}
	case "bottomcenter":
		points = [][3]int{
			{first.GetCenter()[0], first.Position[1] + first.Size[1], 0},
			{first.GetCenter()[0], second.Position[1] - (second.Position[1] - first.Position[1] - first.Size[1]) / 2, 0},
			{second.GetCenter()[0], second.Position[1] - (second.Position[1] - first.Position[1] - first.Size[1]) / 2, 0},
			{second.GetCenter()[0], second.Position[1], 0},
		}
	case "bottomright":
		points = [][3]int{
			{first.Position[0] + first.Size[0], first.Position[1] + first.Size[1], 0},
			{first.Position[0] + first.Size[0], second.GetCenter()[1], 0},
			{second.Position[0], second.GetCenter()[1], 0},
		}
	}
	return points
}

func (s *Cell) RecalculateSynapses() {
	var synapses = []Synapse{}
	for e := range s.Cells {
		points := generateSynapses(*s, s.Cells[e])
		synapses = append(synapses, Synapse{Points: points, Size: 2, Color: "#fa4372"})
		s.Cells[e].RecalculateSynapses()
	}
	s.Synapses = synapses
}

type CellData struct {
	Errors []core.ErrorMsg `json:"errors"`
	Data   Cell            `json:"data"`
}

func (u *CellData) Read(r *http.Response) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(body), &u)
	defer r.Body.Close()
}

func GetRelativePosition(first, second Cell) string {
	res := "undefined"
	sc := second.GetCenter()
	if sc[0] < first.Position[0] && sc[1] < first.Position[1] {
		res = "topleft"
	} else if sc[0] > first.Position[0] && sc[0] < first.Position[0]+first.Size[0] && sc[1] < first.Position[1] {
		res = "topcenter"
	} else if sc[0] > first.Position[0]+first.Size[0] && sc[1] < first.Position[1] {
		res = "topright"
	} else if sc[0] < first.Position[0] && sc[1] > first.Position[1] && sc[1] < first.Position[1]+first.Size[1] {
		res = "midleft"
	} else if sc[0] > first.Position[0]+first.Size[0] && sc[1] > first.Position[1] && sc[1] < first.Position[1]+first.Size[1] {
		res = "midright"
	} else if sc[0] < first.Position[0] && sc[1] > first.Position[1]+first.Size[1] {
		res = "bottomleft"
	} else if sc[0] > first.Position[0] && sc[0] < first.Position[0]+first.Size[0] && sc[1] > first.Position[1]+first.Size[1] {
		res = "bottomcenter"
	} else if sc[0] > first.Position[0]+first.Size[0] && sc[1] > first.Position[1]+first.Size[1] {
		res = "bottomright"
	}
	return res
}

func IsCollision(first, second Cell) bool {
	if second.Position[0] > first.Position[0]+first.Size[0] ||
		first.Position[0] > second.Position[0]+second.Size[0] ||
		second.Position[1] > first.Position[1]+first.Size[1] ||
		first.Position[1] > second.Position[1]+second.Size[1] {
		return false
	} else {
		return true
	}
}
