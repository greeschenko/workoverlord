package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var (
	Dbpath = os.Getenv("HOME") + "/prodev/"
	Dbfile = Dbpath + "MIND2"
)

// initialyze file db
func initDb() {
	if _, err := os.Stat(Dbpath); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(Dbpath, os.ModePerm)
		fmt.Println("- db path created")
	}

	if _, err := os.Stat(Dbfile); errors.Is(err, os.ErrNotExist) {
		saveData()
		fmt.Println("===============================")
		fmt.Println("New MIND created")
		fmt.Println("Run app againe to continue")
		os.Exit(1)
	} else {
		tmpdata, err := os.ReadFile(Dbfile)
		checkErr(err)
		json.Unmarshal(tmpdata, &USERMIND)
	}
}

func saveData() {
	file, err := os.Create(Dbfile)
	checkErr(err)
	defer file.Close()

	USERMINDjson, err := json.MarshalIndent(USERMIND, " ", " ")
	checkErr(err)

	w := bufio.NewWriter(file)
	n4, _ := w.Write(USERMINDjson)
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}
