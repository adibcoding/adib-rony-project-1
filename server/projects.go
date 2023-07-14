package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



type Project struct {
  gorm.Model
  Title  string
  LinkUrl string
}

func AllProjects(w http.ResponseWriter, r *http.Request) {
  // fmt.Fprintf(w, "All Projects")
  db := connectionDB()
  var projects []Project
  db.Find(&projects)
	fmt.Println("{}", projects)

	json.NewEncoder(w).Encode(projects)
}

func AddProjects(w http.ResponseWriter, r *http.Request) {
  // fmt.Fprintf(w, "Add Projects")
  db := connectionDB()
  err := r.ParseForm()
  if err != nil {
      http.Error(w, "Failed to parse request body", http.StatusBadRequest)
      return
  }

  // Access form data
  formData := r.Form
  // Access specific form fields
  title := formData.Get("title")
  linkUrl := formData.Get("linkUrl")

  // Use the form data as needed
  fmt.Println("title:", title)
  fmt.Println("linkUrl:", linkUrl)

	db.Create(&Project{Title: title, LinkUrl: linkUrl})
	fmt.Fprintf(w, "New User Successfully Created")
}

func connectionDB() *gorm.DB{
  dsn := "host=localhost port=5432 user=postgres password=postgres dbname=projects sslmode=disable"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  return db
}

func InitialMigration(){
  dsn := "host=localhost port=5432 user=postgres password=postgres dbname=projects sslmode=disable"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(&Project{})

  // Create
  db.Create(&Project{Title: "D42", LinkUrl: "testing"})
}


