package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-rest-framework/core"
)

type Mind []Cell

type MindData struct {
	Errors []core.ErrorMsg `json:"errors"`
	Data   Mind            `json:"data"`
}

func (u *MindData) Read(r *http.Response) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(body), &u)
	defer r.Body.Close()
}

func handleListInRecursiveChanels(data []Cell, name, mach string, fool bool, c chan *Cell) {
	for k, s := range data {
		if s.getValue(name) == mach {
			c <- &data[k]
			return
		}
		go handleListInRecursiveChanels(s.Cells, name, mach, fool, c)
	}
}

func (m Mind) Find(name, mach string, fool bool) *Cell {
	c := make(chan *Cell)

	go handleListInRecursiveChanels(m, name, mach, fool, c)

	res := <-c

	return res
}

func (m *Mind) Extend(s Cell, parentid string) Cell {
	s.genID(parentid)
	if parentid == "0" {
		*m = append(*m, s)
	} else {
		t := m.Find("id", parentid, true)
		if t != nil {
			t.Cells = append(t.Cells, s)
		} else {
			log.Println("not found " + parentid)
		}
	}

	return s
}
