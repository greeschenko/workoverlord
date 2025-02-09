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
func initDb() error {
	if _, err := os.Stat(Dbpath); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(Dbpath, os.ModePerm)
		fmt.Println("- db path created")
	}

	if _, err := os.Stat(Dbfile); errors.Is(err, os.ErrNotExist) {
		fmt.Println("===============================")
		fmt.Println("New MIND created")
	} else {
		tmpdata, err := os.ReadFile(Dbfile)
        if err != nil {
            return err
        }
        data, err := DataDescript(tmpdata)
        if err != nil {
            return err
        }
		json.Unmarshal(data, &USERMIND)
	}

    return nil
}

func saveData() {
	file, err := os.Create(Dbfile)
	checkErr(err)
	defer file.Close()

	USERMINDjson, err := json.MarshalIndent(USERMIND, " ", " ")
	checkErr(err)

	USERMINDjsonSecret := DataEnctypt(USERMINDjson)

	w := bufio.NewWriter(file)
	n4, _ := w.Write(USERMINDjsonSecret)
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}
