package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/mail"
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
	UserEmail string
	UserPass  string

	Dbpath   = os.Getenv("HOME") + "/prodev2/"
	Dbfile   = Dbpath + "MIND"
	frontend []byte
)

type Config struct {
	UserEmail     string `json:"useremail"`
	Usersecretkey string `json:"usersecretkey"`
}

func askString(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, err := reader.ReadString('\n')
	return strings.TrimSpace(text), err
}

func askPass(prompt string) (string, error) {
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	password := string(bytePassword)
	return strings.TrimSpace(password), nil
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func init() {
	//fmt.Println(logo)
	if _, err := os.Stat(Dbpath); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(Dbpath, os.ModePerm)
		fmt.Println("- db path created")
	}

	if _, err := os.Stat(Dbfile); errors.Is(err, os.ErrNotExist) {

		for !isEmailValid(UserEmail) {
			if UserEmail != "" {
				fmt.Println("[ERR] Email has incorect format, try again please")
			}
			UserEmail, _ = askString("Enter email: ")
		}

		tmpuserpass := ""
		tmpuserpassre := " "

		for tmpuserpass != tmpuserpassre {
			tmpuserpass, _ = askPass("\n Enter new password: ")
			tmpuserpassre, _ = askPass("\n Enter new password re: ")
		}

		SECRETKEY = sha256.Sum256([]byte(tmpuserpass))

		config := Config{
			UserEmail:     UserEmail,
			Usersecretkey: fmt.Sprintf("%x", SECRETKEY),
		}

		configString, _ := json.Marshal(config)

		configCell := Cell{
			Data:   string(configString),
			Status: "config",
		}

		USERMIND = append(USERMIND, configCell)

		saveData()

		fmt.Println("===============================")
		fmt.Println("New MAIND created")
		fmt.Println("Run app againe to continue")
		os.Exit(1)
	} else {

		UserPass, _ = askPass("Enter Password: ")
		//fmt.Printf("Password: %s\n", UserPass)
		//fmt.Print(string(dat))

		SECRETKEY = sha256.Sum256([]byte(UserPass))

		fmt.Println("SECRETKEY: ", SECRETKEY, len(SECRETKEY))

		tmpdata, err := os.ReadFile(Dbfile)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(DataDescript(tmpdata), &USERMIND)

		//compile frontend
		index, _ := os.ReadFile("./client/build/index.html")
		precss := []byte("<style type='text/css'>")
		css, _ := os.ReadFile("./client/build/static/css/main.css")
		postcss := []byte("</style>")
		prejs1 := []byte("<script>")
		js1, _ := os.ReadFile("./client/build/static/js/chunk.js")
		postjs1 := []byte("</script>")
		prejs2 := []byte("<script>")
		js2, _ := os.ReadFile("./client/build/static/js/main.js")
		postjs2 := []byte("</script>")

		frontend = append(frontend, index...)
		frontend = append(frontend, precss...)
		frontend = append(frontend, css...)
		frontend = append(frontend, postcss...)
		frontend = append(frontend, prejs1...)
		frontend = append(frontend, js1...)
		frontend = append(frontend, postjs1...)
		frontend = append(frontend, prejs2...)
		frontend = append(frontend, js2...)
		frontend = append(frontend, postjs2...)
	}
}

func saveData() {
	file, fileerr := os.Create(Dbfile)
	if fileerr != nil {
		log.Fatal(fileerr)
	}
	defer file.Close()

	USERMINDjson, _ := json.MarshalIndent(USERMIND, " ", " ")

	USERMINDjsonSecret := DataEnctypt(USERMINDjson)

	w := bufio.NewWriter(file)
	n4, _ := w.Write(USERMINDjsonSecret)
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}

func logData() {
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, Authorization")
	(*w).Header().Set("Content-Type", "application/json")
}

func main() {

	App.Init()

	App.Config = core.Config{
		MailLogin:    "polonexsender@gmail.com",
		MailPassword: "[Polonex1.emailpass]",
	}

	App.R.HandleFunc("/", actionIndex).Methods("GET")
	App.R.HandleFunc("/cells", actionCellsGetAll).Methods("GET")
	App.R.HandleFunc("/cells/{id}", actionCellsOne).Methods("GET")
	App.R.HandleFunc("/cells/{id}", actionCellsCreate).Methods("POST")
	App.R.HandleFunc("/cells/{id}", actionCellsUpdate).Methods("PATCH")
	App.R.HandleFunc("/cells/{id}", actionCellsDelete).Methods("DELETE")

	go logData()

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
	w.Write(frontend)
}

func actionCellsGetAll(w http.ResponseWriter, r *http.Request) {
	var (
		data Mind
		rsp  = core.Response{Data: &data, Req: r}
	)

	enableCors(&w)

	fmt.Println("USERMIND", USERMIND)

	data = USERMIND

	w.Write(rsp.Make())
}

func actionCellsOne(w http.ResponseWriter, r *http.Request) {
	var (
		data Cell
		rsp  = core.Response{Data: &data, Req: r}
		vars = mux.Vars(r)
	)

	enableCors(&w)

	fmt.Println("USERMIND", USERMIND)

	res := USERMIND.Find("id", vars["id"], true)
	data = *res

	w.Write(rsp.Make())
}

func actionCellsCreate(w http.ResponseWriter, r *http.Request) {
	var (
		model Cell
		rsp   = core.Response{Data: &model, Req: r}
		vars  = mux.Vars(r)
	)

	enableCors(&w)

	if rsp.IsJsonParseDone(r.Body) && rsp.IsValidate() {
		model = USERMIND.Extend(model, vars["id"])
	}

	saveData()

	w.Write(rsp.Make())
}

func actionCellsUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		model Cell
		rsp   = core.Response{Data: &model, Req: r}
		vars  = mux.Vars(r)
	)

	enableCors(&w)

	if rsp.IsJsonParseDone(r.Body) && rsp.IsValidate() {
		USERMIND.Find("id", vars["id"], true).Update(model)
	}

	saveData()

	w.Write(rsp.Make())
}

func actionCellsDelete(w http.ResponseWriter, r *http.Request) {
	var (
		model Cell
		rsp   = core.Response{Data: &model, Req: r}
		vars  = mux.Vars(r)
	)

	enableCors(&w)

	if rsp.IsValidate() {
		USERMIND.DeleteCell(vars["id"])
	}

	saveData()

	w.Write(rsp.Make())
}
