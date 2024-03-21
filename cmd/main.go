package main

import (
	"battle-process/internal/game"
	"battle-process/internal/server"
	"log"
	"net/http"
)

func main() {
	g := game.NewGame()

	srv := server.NewServer(g)

	http.HandleFunc("/connect", srv.ConnectAsPlayer)

	http.HandleFunc("/state", srv.GetGameState)

	http.HandleFunc("/activation/start", srv.StartActivation)
	http.HandleFunc("/activation/stop", srv.StopActivation)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
