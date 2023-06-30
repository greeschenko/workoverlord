package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/go-rest-framework/core"
	"github.com/gorilla/mux"
	"golang.org/x/term"
)

const logo = `
 __        __    _          ___                 _               _
 \ \      / /__ | | ___ __ / _ \__   _____ _ __| | ___  _ __ __| |
  \ \ /\ / / _ \| |/ / '__| | | \ \ / / _ \ '__| |/ _ \| '__/ _' |
   \ V  V / (_) |   <| |  | |_| |\ V /  __/ |  | | (_) | | | (_| |
    \_/\_/ \___/|_|\_\_|   \___/  \_/ \___|_|  |_|\___/|_|  \__,_|
`

var (
	USERMIND  Mind
	SECRETKEY [32]byte
	App       core.App
	UserPass  string

	Dbpath = os.Getenv("HOME") + "/prodev2/"
	Dbfile = Dbpath + "MIND"
)

func askPass() (string, error) {
	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	password := string(bytePassword)
	return strings.TrimSpace(password), nil
}

func init() {
	//fmt.Println(logo)
	if _, err := os.Stat(Dbpath); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(Dbpath, os.ModePerm)
		fmt.Println("- db path created")
	}

	if _, err := os.Stat(Dbfile); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(Dbfile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		fmt.Println("- db file created")
	}

	UserPass, _ = askPass()
	fmt.Printf("Password: %s\n", UserPass)
	//fmt.Print(string(dat))

	SECRETKEY = sha256.Sum256([]byte(UserPass))

	fmt.Println("SECRETKEY: ", SECRETKEY, len(SECRETKEY))

	tmpdata, err := os.ReadFile(Dbfile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DATA: ", tmpdata, len(tmpdata))

	if len(tmpdata) == 0 {
		USERMIND = Mind{}
	}
}

func saveData() {
	n := 0
	tick := time.Tick(15000 * time.Millisecond)
	//boom := time.After(12000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println(USERMIND)
		//case <-boom:
		//fmt.Println("BOOM!")
		//return
		default:
			fmt.Println("    .")
		}
		time.Sleep(15000 * time.Millisecond)
		n++
	}
}

func main() {

	App.Init()

	App.Config = core.Config{
		MailLogin:    "polonexsender@gmail.com",
		MailPassword: "[Polonex1.emailpass]",
	}

	App.R.HandleFunc("/", actionIndex).Methods("GET")
	App.R.HandleFunc("/cells", actionCellsGet).Methods("GET")
	App.R.HandleFunc("/cells/{id}", actionCellsCreate).Methods("POST")
	App.R.HandleFunc("/cells/{id}", actionCellsUpdate).Methods("PATCH")
	//App.R.HandleFunc("/cells/{id}", actionCellsDelete).Methods("DELETE")

	go saveData()

	log.Printf("%s", logo)
	App.Run(":2222")
}

func doRequest(url, proto, userJson, token string) *http.Response {
	reader := strings.NewReader(userJson)
	request, err := http.NewRequest(proto, url, reader)
	if token != "" {
		request.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("REQUEST", url, proto, resp.StatusCode, "RESP HEADER \r\n", resp.Header, "REQ HEADER\r\n", request.Header, userJson)

	return resp
}

func actionIndex(w http.ResponseWriter, r *http.Request) {

	dat, err := os.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(dat)

	//fmt.Println("USERMIND", USERMIND)

	//data = USERMIND

	w.Write([]byte(dat))
}

func actionCellsGet(w http.ResponseWriter, r *http.Request) {
	var (
		data Mind
		rsp  = core.Response{Data: &data, Req: r}
	)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Max-Age", "15")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

	fmt.Println("USERMIND", USERMIND)

	data = USERMIND

	w.Write(rsp.Make())
}

func actionCellsCreate(w http.ResponseWriter, r *http.Request) {
	var (
		model Cell
		rsp   = core.Response{Data: &model, Req: r}
		vars  = mux.Vars(r)
	)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Max-Age", "15")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

	if rsp.IsJsonParseDone(r.Body) && rsp.IsValidate() {
		model = USERMIND.Extend(model, vars["id"])
	}

	w.Write(rsp.Make())
}

func actionCellsUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		model Cell
		rsp   = core.Response{Data: &model, Req: r}
		vars  = mux.Vars(r)
	)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Max-Age", "15")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

	if rsp.IsJsonParseDone(r.Body) && rsp.IsValidate() {
		USERMIND.Find("id", vars["id"], true).Update(&model)
	}

	w.Write(rsp.Make())
}
