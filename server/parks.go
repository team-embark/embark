package main

import (
	"fmt"
	"net/http"
	// "encoding/json"
	firestore "cloud.google.com/go/firestore"
)

func park(w http.ResponseWriter, r *http.Request) {
	token, err := getAuthToken(r.Header.Get("Authorization"))
    if err != nil {
        http.Error(w, "Failed to authenticate user", http.StatusInternalServerError)
        return
    }
    uid := token.UID
	userRef := Client.Collection("users").Doc(uid)
	doc, err := userRef.Get(Ctx)
	if err != nil || !doc.Exists() {
		http.Error(w, "Failed to retrieve context", http.StatusInternalServerError)
		return
	}
	park_name := r.PathValue("park")
	if r.Method == "POST" { // sets park to visted = true
		_, err = userRef.Collection("parks").Doc(park_name).Set(Ctx, map[string]interface{}{
			"visited": true,
		}, firestore.MergeAll)
		if err != nil {
			http.Error(w, "Failed to update park visited status", http.StatusInternalServerError)
		}
	} else if r.Method == "GET" { // return park info
		
	} else {
		fmt.Println("OH NO")
	}
}