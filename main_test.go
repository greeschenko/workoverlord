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
	args    Synapse
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
				Synapse{
					Data:   fake.Words(),
					Status: "new",
					Tags:   "tag1, tag2, tag3",
				},
				[4]int{0, 0, 0, 0},
			},
			{
				"check id gen for parent",
				"ff00",
				Synapse{
					Data:   fake.Words(),
					Status: "new",
					Tags:   "tag1, tag2, tag3",
				},
				[4]int{0, 0, 0, 0},
			},
		},
		"createroot": {
			//create two synapse in root
			{
				"create synapse #1",
				"POST",
				Synapse{
					Data:   fake.Words(),
					Status: "new",
					Tags:   "tag1, tag2, tag3",
				},
				[4]int{0, 0, 0, 0},
			},
			{
				"create synapse #2",
				"POST",
				Synapse{
					Data:   fake.Words(),
					Status: "new",
					Tags:   "tag1, tag2",
				},
				[4]int{0, 0, 0, 0},
			},
		},
		"addchildren": {
			{
				"add synapse #3 to #1",
				"POST",
				Synapse{
					Data:   fake.Words(),
					Status: "new",
					Tags:   "tag1, tag4",
				},
				[4]int{0, 0, 0, 0},
			},
			{
				"add synapse #4 to #1",
				"POST",
				Synapse{
					Data:   fake.Words(),
					Status: "new",
					Tags:   "tag5",
				},
				[4]int{0, 0, 0, 0},
			},
		},
		"update": {
			{
				"update synapse #1",
				"PATCH",
				Synapse{
					Data: fake.Words(),
				},
				[4]int{0, 0, 0, 0},
			},
			{
				"update synapse #1",
				"PATCH",
				Synapse{
					Data: fake.Words(),
					Tags: "tag1",
				},
				[4]int{0, 0, 0, 0},
			},
			{
				"archive synapse #1",
				"PATCH",
				Synapse{
					Status: "archived",
				},
				[4]int{0, 0, 0, 0},
			},
			{
				"delete synapse #1",
				"PATCH",
				Synapse{
					Status: "deleted",
				},
				[4]int{0, 0, 0, 0},
			},
		},
		"index": {
			{
				"get all data",
				"GET",
				Synapse{},
				[4]int{0, 0, 0, 0},
			},
		},
		"move": {
			{
				"move synapse #3 to #2",
				"PATCH",
				Synapse{},
				[4]int{0, 0, 1, 0}, //move from position 0 0 to position 1 0
			},
			{
				"move synapse #4 to #2",
				"PATCH",
				Synapse{},
				[4]int{0, 0, 1, 0}, //repeat move from 0 0 to 1 0
			},
		},
		"changeorder": {
			{
				"change sub element order in #2",
				"PATCH",
				Synapse{},
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
	var url = Apiurl + "/synapse"
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

			tmpsynapse := tt.args

			tmpsynapse.genID(tt.proto)

			if tmpsynapse.ID == "" {
				t.Errorf("ID is empty, gen dont work")
			}

			if len(tmpsynapse.ID) != 4 && len(tmpsynapse.ID) != 9 {
				t.Errorf("ID has wrong len = %d, want %v", len(tmpsynapse.ID), 4)
			}
		})
	}
}

// Test create actions
func Test_actionsCreateRoot(t *testing.T) {
	var url = Apiurl + "/synapse"
	tests := TestData["createroot"]
	//get test user auntification token
	for k, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var synapse SynapseData
			iurl := url + "/0"

			userJson, err := json.MarshalIndent(tt.args, " ", " ")
			if err != nil {
				log.Fatal(err)
			}

			resp := doRequest(iurl, tt.proto, string(userJson), "")

			if resp.StatusCode != 200 {
				t.Errorf("Wrong Response status = %s, want %v", resp.Status, 200)
			}

			synapse.Read(resp)

			if synapse.Data.ID == "" {
				t.Errorf("Wrong ID synapse not created")
			}
			TestData["createroot"][k].args.ID = synapse.Data.ID
		})
	}
}

func Test_actionsCreateChildren(t *testing.T) {
	var url = Apiurl + "/synapse"
	tests := TestData["addchildren"]
	//get test user auntification token
	for k, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var synapse SynapseData
			iurl := url + "/" + TestData["createroot"][0].args.ID

			userJson, err := json.MarshalIndent(tt.args, " ", " ")
			if err != nil {
				log.Fatal(err)
			}

			resp := doRequest(iurl, tt.proto, string(userJson), "")

			if resp.StatusCode != 200 {
				t.Errorf("Wrong Response status = %s, want %v", resp.Status, 200)
			}

			synapse.Read(resp)

			if synapse.Data.ID == "" {
				t.Errorf("Wrong ID synapse not created")
			}
			TestData["addchildren"][k].args.ID = synapse.Data.ID
		})
	}
}

func Test_actionsUpdate(t *testing.T) {
	var url = Apiurl + "/synapse"
	tests := TestData["update"]
	//get test user auntification token
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var synapse SynapseData
			iurl := url + "/" + TestData["createroot"][0].args.ID

			userJson, err := json.MarshalIndent(tt.args, " ", " ")
			if err != nil {
				log.Fatal(err)
			}

			resp := doRequest(iurl, tt.proto, string(userJson), "")

			if resp.StatusCode != 200 {
				t.Errorf("Wrong Response status = %s, want %v", resp.Status, 200)
			}

			synapse.Read(resp)

			if len(synapse.Errors) > 0 {
				t.Error("Api return Errs", synapse.Errors)
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
