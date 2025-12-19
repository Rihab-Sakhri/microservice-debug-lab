package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Rihab"},
	{ID: 2, Name: "Ahmed"},
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// Intentional bug: returning nil instead of users slice in some cases
	if r.URL.Query().Get("bug") == "true" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong!"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/users", getUsers)
	log.Println("User service running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	{
		fmt.Println("User service running on :8080")
		notifyFlask(1)
		select {}
	}
}

type Notification struct {
	ID int `json:"id"`
}

func notifyFlask(userID int) {
	url := "http://notify-service-python:5000/notify" // use Docker service name

	notification := Notification{ID: userID}
	jsonData, _ := json.Marshal(notification)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending notification:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Flask response status:", resp.Status)
}
