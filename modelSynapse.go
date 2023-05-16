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

type Synapse struct {
	ID       string    `json:"id"`
	Data     string    `json:"data"`
	Status   string    `json:"status"`
	Tags     string    `json:"tags,omitempty"`
	Size     [2]int    `json:"size,omitempty"`
	Position [2]int    `json:"position,omitempty"`
	Routs    [][2]int  `json:"routs,omitempty"`
	Pointers [][2]int  `json:"pointers,omitempty"`
	Synapses []Synapse `json:"synapses"`
}

func (s Synapse) getValue(name string) string {
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

func (s *Synapse) genID(parentid string) {
	resid, _ := randomHex(2)
	if parentid != "0" {
		s.ID = parentid + " " + resid
	} else {
		s.ID = resid
	}
	fmt.Println("GENERATED ID", s.ID)
}

func (s *Synapse) Update(newdata *Synapse) {
	s = newdata
}

type SynapseData struct {
	Errors []core.ErrorMsg `json:"errors"`
	Data   Synapse         `json:"data"`
}

func (u *SynapseData) Read(r *http.Response) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(body), &u)
	defer r.Body.Close()
}
