package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
    "fmt"
    "strings"

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

	fmt.Println("OOOOOOO", idlistres)
	fmt.Println("IDDDDDD", id)

	tmp = deleteCellListHandle(tmp, id, idlistres, 0)

	*m = tmp
}

func deleteCellListHandle(cc []Cell, id string, idlistres []string, idindex int) []Cell {
    fmt.Println("EEEEEEEEEEEEEE", cc, id, idlistres, idindex)
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
                fmt.Println("RRRRRRRR", cc[j].Cells)
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

            var newpoints [][3]int
            var parentcellright [3]int
            var parentcellbottom [3]int
            var parentcellleft [3]int
            var parentcelltop [3]int
            var newcellright [3]int
            var newcellbottom [3]int
            var newcellleft [3]int
            var newcelltop [3]int

            parentcellright[0] = parentcell.Position[0] + parentcell.Size[0]
            parentcellright[1] = parentcell.Position[1] + parentcell.Size[1] / 2
            parentcellbottom[0] = parentcell.Position[0] + parentcell.Size[0] / 2
            parentcellbottom[1] = parentcell.Position[1] + parentcell.Size[1]
            parentcellleft[0] = parentcell.Position[0]
            parentcellleft[1] = parentcell.Position[1] + parentcell.Size[1] / 2
            parentcelltop[0] = parentcell.Position[0] + parentcell.Size[0] / 2
            parentcelltop[1] = parentcell.Position[1]

            newcellright[0] = newcell.Position[0] + newcell.Size[0]
            newcellright[1] = newcell.Position[1] + newcell.Size[1] / 2
            newcellbottom[0] = newcell.Position[0] + newcell.Size[0] / 2
            newcellbottom[1] = newcell.Position[1] + newcell.Size[1]
            newcellleft[0] = newcell.Position[0]
            newcellleft[1] = newcell.Position[1] + newcell.Size[1] / 2
            newcelltop[0] = newcell.Position[0] + newcell.Size[0] / 2
            newcelltop[1] = newcell.Position[1]

            if parentcellright[0] < newcellleft[0] {
                //element on right
                newpoints = append(newpoints, [3]int{newcellleft[0], newcellleft[1]}, [3]int{parentcellright[0], parentcellright[1]})
            }else if parentcellbottom[1] < newcelltop[1] {
                //element on bottom
                newpoints = append(newpoints, [3]int{newcelltop[0], newcelltop[1]}, [3]int{parentcellbottom[0], parentcellbottom[1]})
            }else if parentcellleft[0] > newcellright[0] {
                //element on left
                newpoints = append(newpoints, [3]int{newcellright[0], newcellright[1]}, [3]int{parentcellleft[0], parentcellleft[1]})
            }else if parentcelltop[1] > newcellbottom[1] {
                //element on top
                newpoints = append(newpoints, [3]int{newcellbottom[0], newcellbottom[1]}, [3]int{parentcelltop[0], parentcelltop[1]})
            }

            if len(newpoints) > 0 {
                newcell.Synapses = []Synapse{{
                    Points: newpoints,
                    Size: 1,
                    Color: "red",
                }}
            }

			parentcell.Cells = append(parentcell.Cells, newcell)
		} else {
			log.Println("not found " + parentid)
		}
	}

	return newcell
}
