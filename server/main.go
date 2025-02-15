package main

import (
	"context"
	"fmt"
	"net/http"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/rs/cors"
	"google.golang.org/api/option"
)

var Ctx context.Context = context.Background()
var Opt option.ClientOption = option.WithCredentialsFile("embark.json")
var (
	App        *firebase.App
	AuthClient *auth.Client
	Client     *firestore.Client
)

func main() {
	mux := &http.ServeMux{}

	mux.HandleFunc("/", helloWorld)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Context-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(mux)

    go initApps()

	http.ListenAndServe(":8080", corsHandler)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
