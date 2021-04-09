package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/m3talsmith/talman/game"
)

func main() {
	initConfig() // loads config from config.go

	http.HandleFunc("/", runGame)

	fmt.Println("Loading Talman on localhost:3000")
	http.ListenAndServe(url, nil)
}

func runGame(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "time", time.Now())
	game.Next(ctx, w)
}
