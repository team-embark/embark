package main

import (
	"fmt"
	firebase "firebase.google.com/go/v4"
)

func initApps() {
	var err error
	App, err = firebase.NewApp(Ctx, nil, Opt)
	if err != nil {
		fmt.Println("Failed to initiialize firebase app:", err)
		return
	}
	AuthClient, err = App.Auth(Ctx)
	if err != nil {
		fmt.Println("Failed to initiialize firebase auth client:", err)
		return
	}
	Client, err = App.Firestore(Ctx)
	if err != nil {
		fmt.Println("Failed to initialize firestore client:", err)
		return
	}
}
