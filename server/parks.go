package main

import (
	"net/http"
	"encoding/json"
	firestore "cloud.google.com/go/firestore"
)

func visitPark(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Invalid response method", http.StatusBadRequest)
	}
}

func getAllParkInfo(w http.ResponseWriter, r *http.Request) { // returns true/false for each park user has visited or not
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

	// return parks' info
	parks := [3]string{"yosemite", "josh", "redwood"}
	var data map[string]interface{}
	for _,p := range parks {// add p into data
		// p : true/false
		pdata, err := userRef.Collection("parks").Doc(p).Get(Ctx)
		if err != nil {
			http.Error(w, "poop", http.StatusInternalServerError)
		}
		m := pdata.Data()
		data[p] = m["visited"]
	}
	jsonResp, err := json.Marshal(data)
    if err != nil {
        http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResp)
}