package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var db *sql.DB

func ConnectToDatabase() *sql.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var dbUrl = os.Getenv("DATABASE_URL")
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbUrl, err)
		os.Exit(1)
	}
	fmt.Println("Connected to database")
	return db
}

func RenderHTMLTemplate(w http.ResponseWriter, data interface{}, templatePath ...string) {
	t, err := template.ParseFiles(templatePath...)
	if err != nil {
		http.Error(w, "Failed to parse templates", http.StatusInternalServerError)
		log.Println("Error parsing templates:", err)
		return
	}
	log.Println("Rendering template:", templatePath, "with data:", data)
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Println("Error rendering template:", err)
	}
}
