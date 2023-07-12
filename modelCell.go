package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-rest-framework/core"
)

type Cell struct {
	ID       string    `json:"id"`
	Data     string    `json:"data"`
	Status   string    `json:"status"`
	Tags     string    `json:"tags,omitempty"`
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
	case "tags":
		res = s.Tags
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

func (s *Cell) Update(newdata *Cell) {
	s = newdata
}

func (s *Cell) Delete() {
	s = nil
}

type CellData struct {
	Errors []core.ErrorMsg `json:"errors"`
	Data   Cell            `json:"data"`
}

func (u *CellData) Read(r *http.Response) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(body), &u)
	defer r.Body.Close()
}
