package main

import (
    "encoding/json"
    "log"
    "net/http"
    "database/sql"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
		"io/ioutil"
)

var db *sql.DB
var err error

func show(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
   results, err := db.Query("SELECT * FROM test")
    if err != nil {
        panic(err.Error())
    }
    var printStr []string
    for results.Next(){
        var id int
        var str string
        err = results.Scan(&id, &str)
        if err != nil {
            panic(err.Error())
        }
        printStr = append(printStr, str)
        log.Println(str)
    }
    json.NewEncoder(w).Encode(printStr)
}


func add(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var body, err= ioutil.ReadAll(r.Body)
if err != nil {
        panic(err.Error())
    }
log.Println("Name received: ",string(body))
var nm=string(body)
var id=1
	insert,err := db.Prepare("INSERT INTO test(id, str) VALUES(?, ?)")
        if err != nil {
            panic(err.Error())
        }
        insert.Exec(id,nm)

    json.NewEncoder(w).Encode("Added")
}

func del(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
/*
	insert,err := db.Prepare("INSERT INTO test(id, str) VALUES(?, ?)")
        if err != nil {
            panic(err.Error())
        }
        insert.Exec(id,nm)
*/
		results, err := db.Prepare("DELETE FROM test")
		if err != nil {
            panic(err.Error())
        }
        results.Exec()
    json.NewEncoder(w).Encode("All Deleted")
}

func handleRequests() {
    router := mux.NewRouter()
    router.HandleFunc("/show", show)
    router.HandleFunc("/add", add)
		router.HandleFunc("/del", del)
    log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
log.Println("Starting the server...")
 /*   db,err= sql.Open("mysql","ketan:ketan@tcp(mysqldb:3306)/name")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    _,err = db.Exec("CREATE TABLE IF NOT EXISTS test ( id integer, str varchar(500) )")
    if err != nil {
       panic(err)
    }
log.Println("Connection to db established")
*/
    handleRequests()
}
