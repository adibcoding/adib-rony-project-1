package main
import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func handleRequests() {
  myRouter := mux.NewRouter().StrictSlash(true)
  myRouter.HandleFunc("/projects", AllProjects).Methods("GET")
  myRouter.HandleFunc("/projects", AddProjects).Methods("POST")
  log.Fatal(http.ListenAndServe(":8081", myRouter))
}


func main() {
  fmt.Println("Go ORM Tutorial")
  // InitialMigration()
  // Handle Subsequent requests
  handleRequests()
}