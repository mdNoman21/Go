package main

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

const baseURL = "http://localhost:8080"

// func urlShortener(url string) {

// }
func generateRandomCharacter(pool string) (string, error) {
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(pool))))
	if err != nil {
		return "", err
	}
	return string(pool[index.Int64()]), nil
}

type ShortId string

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, longURL, 301)
}

var longURL string

func main() {
	// Capturing connection properties
	urlMap := make(map[string]string)
	cfg := mysql.Config{
		User:   "root",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "url_schema",
	}
	// Get a database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	characterPool := "1234567890QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm"
	fmt.Println("Enter URL:")
	fmt.Scanln(&longURL)
	shortId := ShortId("")
	for i := 0; i < 10; i++ {
		shortURL := baseURL
		char, err := generateRandomCharacter(characterPool)
		if err != nil {
			fmt.Println("Error generating random character:", err)
			return
		}
		shortId += ShortId(char)
		shortURL += "/" + string(shortId)
		urlMap[longURL] = shortURL
	}
	finalURL := urlMap[longURL]
	_, err = db.Exec("Insert into URLs (long_url,short_url) VALUES (?, ?)", longURL, finalURL)
	if err != nil {
		fmt.Println("Error inserting data into the database:", err)
	}

	defer db.Close()

	http.HandleFunc(finalURL, redirect)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
