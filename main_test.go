package main

import (
	"encoding/json"
	"fmt"
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
	TestDataReq = map[string][]TestDataItem{
		"checkidgen": {
			{
				"check id gen",
				"0",
				Cell{
					Data:     fake.WordsN(30),
					Status:   "new",
					Size:     [2]int{300, 300},
					Position: [3]int{50, 50, 0},
				},
				[4]int{0, 0, 0, 0},
			},
			{
				"check id gen for parent",
				"ff00",
				Cell{
					Data:     fake.WordsN(30),
					Status:   "new",
					Size:     [2]int{300, 400},
					Position: [3]int{500, 300, 0},
				},
				[4]int{0, 0, 0, 0},
			},
		},
		"createroot": {
			//create two Cell in root
			{
				"check id gen",
				"POST",
				Cell{
					Data:     fake.WordsN(30),
					Status:   "new",
					Size:     [2]int{300, 300},
					Position: [3]int{50, 50, 0},
				},
				[4]int{0, 0, 0, 0},
			},
			{
				"check id gen for parent",
				"POST",
				Cell{
					Data:     fake.WordsN(100),
					Status:   "new",
					Size:     [2]int{300, 400},
					Position: [3]int{750, 300, 0},
				},
				[4]int{0, 0, 0, 0},
			},
		},
		"addchildren": {
			{
				"add Cell #3 to #1",
				"POST",
				Cell{
					Data:     fake.Words() + " https://www.youtube.com/watch?v=T4z-32mXLSY&ab_channel=Nikattica",
					Status:   "new",
					Size:     [2]int{300, 300},
					Position: [3]int{360, 30, 1},
				},
				[4]int{0, 0, 0, 0},
			},
			{
				"add Cell #4 to #1",
				"POST",
				Cell{
					Data:     fake.Words() + " https://cdna.artstation.com/p/assets/images/images/053/956/262/medium/sentron-edgerunner-copy.jpg",
					Status:   "new",
					Size:     [2]int{300, 250},
					Position: [3]int{160, 450, 1},
				},
				[4]int{0, 0, 0, 0},
			},
		},
		"update": {
			{
				"update Cell #1",
				"PATCH",
				Cell{
					Data:     "******************",
					Status:   "new",
					Size:     [2]int{111, 111},
					Position: [3]int{111, 111, 0},
				},
				[4]int{0, 0, 0, 0},
			},
			//			{
			//				"done Cell #1",
			//				"PATCH",
			//				Cell{
			//					Data:     fake.WordsN(30),
			//					Status:   "done",
			//					Size:     [2]int{300, 300},
			//					Position: [3]int{50, 50, 0},
			//				},
			//				[4]int{0, 0, 0, 0},
			//			},
			//			{
			//				"archive Cell #1",
			//				"PATCH",
			//				Cell{
			//					Data:     fake.WordsN(20),
			//					Status:   "archive",
			//					Size:     [2]int{300, 300},
			//					Position: [3]int{50, 50, 0},
			//				},
			//				[4]int{0, 0, 0, 0},
			//			},
			/*
			 *{
			 *  "delete Cell #1",
			 *  "PATCH",
			 *  Cell{
			 *    Status: "deleted",
			 *  },
			 *  [4]int{0, 0, 0, 0},
			 *},
			 */
		},
		//		"index": {
		//			{
		//				"get all data",
		//				"GET",
		//				Cell{},
		//				[4]int{0, 0, 0, 0},
		//			},
		//		},
		//		"move": {
		//			{
		//				"move Cell #3 to #2",
		//				"PATCH",
		//				Cell{},
		//				[4]int{0, 0, 1, 0}, //move from position 0 0 to position 1 0
		//			},
		//			{
		//				"move Cell #4 to #2",
		//				"PATCH",
		//				Cell{},
		//				[4]int{0, 0, 1, 0}, //repeat move from 0 0 to 1 0
		//			},
		//		},
		//		"changeorder": {
		//			{
		//				"change sub element order in #2",
		//				"PATCH",
		//				Cell{},
		//				[4]int{1, 1, 1, 0}, //change position from 1 1 to 1 0
		//			},
		//		},
		//		"delete": {
		//			{
		//				"delete Cell #2",
		//				"DELETE",
		//				Cell{},
		//				[4]int{0, 0, 0, 0},
		//			},
		//			{
		//				"delete Cell #1 1",
		//				"DELETE",
		//				Cell{},
		//				[4]int{0, 0, 0, 0},
		//			},
		//		},
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
	tests := TestDataReq["checkidgen"]
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
	tests := TestDataReq["createroot"]
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
			TestDataReq["createroot"][k].args.ID = Cell.Data.ID
		})
	}
}

func Test_actionsCreateChildren(t *testing.T) {
	var url = Apiurl + "/cells"
	tests := TestDataReq["addchildren"]
	//get test user auntification token
	for k, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var Cell CellData
			iurl := url + "/" + TestDataReq["createroot"][0].args.ID

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
			TestDataReq["addchildren"][k].args.ID = Cell.Data.ID
		})
	}
}

func Test_FIND_and_UPDATE(t *testing.T) {
	TestMind := Mind{
		Cell{ID: "001", Data: "data 1"},
		Cell{ID: "002", Data: "data 2", Cells: []Cell{
			{ID: "002 001", Data: "data 2-1"},
			{ID: "002 002", Data: "data 2-2"},
		}},
		Cell{ID: "003", Data: "data 3"},
	}

	res1 := TestMind.Find("id", "001", true)
	fmt.Println(res1.Data)

	res21 := TestMind.Find("id", "002 001", true)
	fmt.Println(res21.Data)

	TestMind.Find("id", "001", true).Update(Cell{ID: "001", Data: "updated data 1"})

	fmt.Println(TestMind[0].Data)

	TestMind.Find("id", "002 001", true).Update(Cell{ID: "002 001", Data: "updated data 2 - 1"})

	fmt.Println(TestMind[1].Cells[0].Data)
	fmt.Println(TestMind)
}

func Test_actionsUpdate(t *testing.T) {
	var url = Apiurl + "/cells"
	tests := TestDataReq["update"]
	//get test user auntification token
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var Cell CellData
			iurl := url + "/" + TestDataReq["createroot"][0].args.ID

      tt.args.ID = TestDataReq["createroot"][0].args.ID

			userJson, err := json.MarshalIndent(tt.args, " ", " ")
			if err != nil {
				log.Fatal(err)
			}

			resp := doRequest(iurl, tt.proto, string(userJson), "")

			if resp.StatusCode != 200 {
				t.Errorf("Wrong Response status = %s, want %v", resp.Status, 200)
			}

			Cell.Read(resp)

			//fmt.Println("CELL DATA AFTER UPDATE", Cell)

			if len(Cell.Errors) > 0 {
				t.Error("Api return Errs", Cell.Errors)
			}

//			var Cell2 CellData
//
//			resp1 := doRequest(url+"/"+TestDataReq["createroot"][0].args.ID, "GET", "", "")
//
//			if resp1.StatusCode != 200 {
//				t.Errorf("Wrong Check Response status = %s, want %v", resp1.Status, 200)
//			}
//
//			Cell2.Read(resp1)
//
//			if Cell.Data.Data != Cell2.Data.Data {
//				t.Errorf("Wrong Updated Data = %s, want %s", Cell.Data.Data, Cell2.Data.Data)
//			}
//
//			if Cell.Data.Status != Cell2.Data.Status {
//				t.Errorf("Wrong Updated Status = %s, want %s", Cell.Data.Status, Cell2.Data.Status)
//			}
		})
	}
}

/*
 *func Test_actionsIndex(t *testing.T) {
 *  var url = Apiurl + "/"
 *  tests := TestData["index"]
 *  //get test user auntification token
 *  for _, tt := range tests {
 *    t.Run(tt.name, func(t *testing.T) {
 *      var mind MindData
 *      iurl := url
 *      resp := doRequest(iurl, tt.proto, "", "")
 *
 *      if resp.StatusCode != 200 {
 *        t.Errorf("Wrong Response status = %s, want %v", resp.Status, 200)
 *      }
 *
 *      mind.Read(resp)
 *
 *      if len(mind.Data) == 0 {
 *        t.Errorf("Wrong element count = 0, want > 0")
 *      }
 *
 *      if len(mind.Errors) > 0 {
 *        t.Error("Api return Errs", mind.Errors)
 *      }
 *
 *      TmpTestMind = mind.Data
 *    })
 *  }
 *}
 */

func Test_actionsDelete(t *testing.T) {
	var url = Apiurl + "/cells"
	tests := TestDataReq["delete"]
	//get test user auntification token
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mind MindData
			var iurl string
			if tt.name == "delete Cell #2" {
				iurl = url + "/" + TestDataReq["createroot"][1].args.ID
			} else if tt.name == "delete Cell #1 1" {
				iurl = url + "/" + TestDataReq["addchildren"][0].args.ID
			}

			resp := doRequest(iurl, tt.proto, "", "")

			if resp.StatusCode != 200 {
				t.Errorf("Wrong Response status = %s, want %v", resp.Status, 200)
			}

			resp1 := doRequest(url, "GET", "", "")

			if resp1.StatusCode != 200 {
				t.Errorf("Wrong Check Response status = %s, want %v", resp.Status, 200)
			}

			mind.Read(resp1)

			fmt.Println("MIND DATA LENGHT", len(mind.Data), TestDataReq["createroot"][0].args.ID)

			if len(mind.Data) != 1 {
				t.Errorf("Wrong element count = %d, want == 1", len(mind.Data))
			}

			if len(mind.Errors) > 0 {
				t.Error("Api return Errs", mind.Errors)
			}

		})
	}
}
