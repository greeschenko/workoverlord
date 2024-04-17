package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-rest-framework/core"
)

type Mind []Cell

type MindData struct {
	Errors []core.ErrorMsg `json:"errors"`
	Data   Mind            `json:"data"`
}

func (u *MindData) Read(r *http.Response) {
	body, err := io.ReadAll(r.Body)
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

func (m *Mind) DeleteCell(id string) {

	idlist := strings.Split(id, " ")
	idlistres := []string{}
	tmpid := ""

	tmp := *m

	for k := range idlist {
		if k == 0 {
			tmpid = idlist[k]
		} else {
			tmpid = tmpid + " " + idlist[k]
		}
		idlistres = append(idlistres, tmpid)
	}

	fmt.Println("IDDDDDD", id)

	tmp = deleteCellListHandle(tmp, id, idlistres, 0)

	*m = tmp
}

func deleteCellListHandle(cc []Cell, id string, idlistres []string, idindex int) []Cell {
	for j := range cc {
		if cc[j].ID == idlistres[idindex] {
			if cc[j].ID == id {
				var tmpres []Cell

				for i := range cc {
					if cc[i].ID != id {
						tmpres = append(tmpres, cc[i])
					}
				}

				cc = tmpres
			} else {
				cc[j].Cells = deleteCellListHandle(cc[j].Cells, id, idlistres, idindex+1)
			}
			break
		}
	}
	return cc
}

func (m *Mind) Extend(newcell Cell, parentid string) Cell {
	newcell.genID(parentid)
	if parentid == "0" {
		*m = append(*m, newcell)
	} else {
		parentcell := m.Find("id", parentid, true)
		if parentcell != nil {
			parentcell.Cells = append(parentcell.Cells, newcell)
		} else {
			log.Println("not found " + parentid)
		}
	}

	return newcell
}

func (m *Mind) RecalculateSynapses() {
    tmpm := *m
    for e := range tmpm {
        tmpm[e].RecalculateSynapses()
    }
}
