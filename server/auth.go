package main

import (
    "fmt"
    "firebase.google.com/go/v4/auth"
    "net/http"
    "strings"
    "time"
    "encoding/json"
)

func getAuthToken(authHeader string) (*auth.Token, error) {
    if authHeader == "" {
        return nil, fmt.Errorf("auth header missing")
    }
    idToken := strings.TrimPrefix(authHeader, "Bearer ")
    if idToken == authHeader {
        return nil, fmt.Errorf("invalid auth token format")
    }
    token, err := AuthClient.VerifyIDToken(Ctx, idToken)
    if err != nil {
        return nil, fmt.Errorf("invalid or expired ID token")
    }
    return token, nil
}

func authUser(w http.ResponseWriter, r *http.Request) {
    token, err := getAuthToken(r.Header.Get("Authorization"))
    if err != nil {
        http.Error(w, "Failed to authenticate user", http.StatusInternalServerError)
        return
    }
    uid := token.UID
    email, _ := token.Claims["email"].(string)

    userRef := Client.Collection("users").Doc(uid)
    doc, err := userRef.Get(Ctx)

    if err != nil {
        http.Error(w, "Failed to get context", http.StatusInternalServerError)
        return
    }

    if !doc.Exists() {
        _, err = userRef.Set(Ctx, map[string]interface{}{
            "uid": uid,
            "email": email,
        })
        if err != nil {
            http.Error(w, "Failed to save user to Firestore", http.StatusInternalServerError)
            return
        }

        _, err = userRef.Collection("parks").Doc("created").Set(Ctx, map[string]interface{}{
            "time": time.Now(),
        })

        if err != nil {
            http.Error(w, "Failed to create park collection", http.StatusInternalServerError)
            return
        }

        parks := [3]string{"yosemite", "josh", "redwood"}
        
        for _,p := range parks {
            _, err = userRef.Collection("parks").Doc(p).Set(Ctx, map[string]interface{}{
                "visited": false,
                "time_visited": nil,
            })
            if err != nil {
                http.Error(w, "Failed to create park", http.StatusInternalServerError)
                return
            }
        }
    }
    resp := map[string]interface{}{
        "uid": uid,
        "email": email,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}
