package main

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/icrowley/fake"
)

type TestDataItem struct {
	name    string
	proto   string
	args    Cell
	indexes [4]int
}

var (
	TmpTestMind Mind
	Apiurl      = "http://localhost:2222"
	TestData    = map[string][]TestDataItem{
		"checkidgen": {
			{
				"check id gen",
				"0",
				Cell{
					Data:     fake.Words(),
					Status:   "new",
					Tags:     "tag1, tag2, tag3",
					Size:     [2]int{100, 100},
					Position: [3]int{0, 0, 0},
				},
				[4]int{0, 0, 0, 0},
			},
			{
				"check id gen for parent",
				"ff00",
				Cell{
					Data:     fake.Words(),
					Status:   "new",
					Tags:     "tag1, tag2, tag3",
					Size:     [2]int{100, 100},
					Position: [3]int{110, 0, 0},
				},
				[4]int{0, 0, 0, 0},
			},
		},
		"createroot": {
			//create two Cell in root
			{
				"create Cell #1",
				"POST",
				Cell{
					Data:   fake.Words(),
					Status: "new",
					Tags:   "tag1, tag2, tag3",
					Size:     [2]int{100, 100},
					Position: [3]int{0, 0, 0},
				},
				[4]int{0, 0, 0, 0},
			},
			{
				"create Cell #2",
				"POST",
				Cell{
					Data:   fake.Words(),
					Status: "new",
					Tags:   "tag1, tag2",
					Size:     [2]int{100, 100},
					Position: [3]int{110, 0, 0},
				},
				[4]int{0, 0, 0, 0},
			},
		},
		"addchildren": {
			{
				"add Cell #3 to #1",
				"POST",
				Cell{
					Data:   fake.Words(),
					Status: "new",
					Tags:   "tag1, tag4",
					Size:     [2]int{100, 100},
					Position: [3]int{160, 50, 1},
				},
				[4]int{0, 0, 0, 0},
			},
			{
				"add Cell #4 to #1",
				"POST",
				Cell{
					Data:   fake.Words(),
					Status: "new",
					Tags:   "tag5",
					Size:     [2]int{100, 100},
					Position: [3]int{160, 150, 1},
				},
				[4]int{0, 0, 0, 0},
			},
		},
		"update": {
			{
				"update Cell #1",
				"PATCH",
				Cell{
					Data: fake.Words(),
				},
				[4]int{0, 0, 0, 0},
			},
			{
				"update Cell #1",
				"PATCH",
				Cell{
					Data: fake.Words(),
					Tags: "tag1",
				},
				[4]int{0, 0, 0, 0},
			},
			{
				"archive Cell #1",
				"PATCH",
				Cell{
					Status: "archived",
				},
				[4]int{0, 0, 0, 0},
			},
			{
				"delete Cell #1",
				"PATCH",
				Cell{
					Status: "deleted",
				},
				[4]int{0, 0, 0, 0},
			},
		},
		"index": {
			{
				"get all data",
				"GET",
				Cell{},
				[4]int{0, 0, 0, 0},
			},
		},
		"move": {
			{
				"move Cell #3 to #2",
				"PATCH",
				Cell{},
				[4]int{0, 0, 1, 0}, //move from position 0 0 to position 1 0
			},
			{
				"move Cell #4 to #2",
				"PATCH",
				Cell{},
				[4]int{0, 0, 1, 0}, //repeat move from 0 0 to 1 0
			},
		},
		"changeorder": {
			{
				"change sub element order in #2",
				"PATCH",
				Cell{},
				[4]int{1, 1, 1, 0}, //change position from 1 1 to 1 0
			},
		},
	}
)

//TODO decide save all maind or each element in single
// save all in one object
//    load object on frontend as json
//    actions
//        create
//        update
//  //        move
//  //        change order
//TODO
//
//CRUD
// Create Task /workoverlord/task POST
// Get Task /workoverlord/task/id GET
// Update Task /workoverlord/task/id PATCH
// Delete Task /workoverlord/task/id DELETE

// Test actions without auntification
/*func Test_actionsWithoutAuntification(t *testing.T) {
	var url = Apiurl + "/cells"
	tests := TestData["auntification"]
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ID string
			if tt.name != "create" {
				ID = "/notexistingid"
			}

			resp := doRequest(url+ID, tt.proto, "[]", "")
			if resp.StatusCode != 401 {
				t.Errorf("Wrong Response status = %s, want %v", resp.Status, 401)
			}
		})
	}
}*/

// Test create actions
func Test_GenID(t *testing.T) {
	tests := TestData["checkidgen"]
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tmpCell := tt.args

			tmpCell.genID(tt.proto)

			if tmpCell.ID == "" {
				t.Errorf("ID is empty, gen dont work")
			}

			if len(tmpCell.ID) != 4 && len(tmpCell.ID) != 9 {
				t.Errorf("ID has wrong len = %d, want %v", len(tmpCell.ID), 4)
			}
		})
	}
}

// Test create actions
func Test_actionsCreateRoot(t *testing.T) {
	var url = Apiurl + "/cells"
	tests := TestData["createroot"]
	//get test user auntification token
	for k, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var Cell CellData
			iurl := url + "/0"

			userJson, err := json.MarshalIndent(tt.args, " ", " ")
			if err != nil {
				log.Fatal(err)
			}

			resp := doRequest(iurl, tt.proto, string(userJson), "")

			if resp.StatusCode != 200 {
				t.Errorf("Wrong Response status = %s, want %v", resp.Status, 200)
			}

			Cell.Read(resp)

			if Cell.Data.ID == "" {
				t.Errorf("Wrong ID Cell not created")
			}
			TestData["createroot"][k].args.ID = Cell.Data.ID
		})
	}
}

func Test_actionsCreateChildren(t *testing.T) {
	var url = Apiurl + "/cells"
	tests := TestData["addchildren"]
	//get test user auntification token
	for k, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var Cell CellData
			iurl := url + "/" + TestData["createroot"][0].args.ID

			userJson, err := json.MarshalIndent(tt.args, " ", " ")
			if err != nil {
				log.Fatal(err)
			}

			resp := doRequest(iurl, tt.proto, string(userJson), "")

			if resp.StatusCode != 200 {
				t.Errorf("Wrong Response status = %s, want %v", resp.Status, 200)
			}

			Cell.Read(resp)

			if Cell.Data.ID == "" {
				t.Errorf("Wrong ID Cell not created")
			}
			TestData["addchildren"][k].args.ID = Cell.Data.ID
		})
	}
}

func Test_actionsUpdate(t *testing.T) {
	var url = Apiurl + "/cells"
	tests := TestData["update"]
	//get test user auntification token
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var Cell CellData
			iurl := url + "/" + TestData["createroot"][0].args.ID

			userJson, err := json.MarshalIndent(tt.args, " ", " ")
			if err != nil {
				log.Fatal(err)
			}

			resp := doRequest(iurl, tt.proto, string(userJson), "")

			if resp.StatusCode != 200 {
				t.Errorf("Wrong Response status = %s, want %v", resp.Status, 200)
			}

			Cell.Read(resp)

			if len(Cell.Errors) > 0 {
				t.Error("Api return Errs", Cell.Errors)
			}
		})
	}
}

func Test_actionsIndex(t *testing.T) {
	var url = Apiurl + "/"
	tests := TestData["index"]
	//get test user auntification token
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mind MindData
			iurl := url
			resp := doRequest(iurl, tt.proto, "", "")

			if resp.StatusCode != 200 {
				t.Errorf("Wrong Response status = %s, want %v", resp.Status, 200)
			}

			mind.Read(resp)

			if len(mind.Data) == 0 {
				t.Errorf("Wrong element count = 0, want > 0")
			}

			if len(mind.Errors) > 0 {
				t.Error("Api return Errs", mind.Errors)
			}

			TmpTestMind = mind.Data
		})
	}
}
