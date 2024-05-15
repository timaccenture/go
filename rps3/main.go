package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"slices"
)

var moves = map[string]string{
	"1": "rock",
	"2": "paper",
	"3": "scissor",
}

var userMinusComp = map[string][]int{
	"Client won": {1, 2, -3},
	"Server won": {-1, -2, 3},
}

func generateMoveAndEvaluate(move string) (string, string) {
	clientMove := move
	validInputs := []string{"rock", "paper", "scissor"}
	randNumber := rand.IntN(len(validInputs))
	serverMove := validInputs[randNumber]

	diff := len(clientMove) - len(serverMove)

	if diff == 0 {
		return "Server and client tied", serverMove
	} else if slices.Contains(userMinusComp["Client won"], diff) {
		return "The Client won.", serverMove
	} else if slices.Contains(userMinusComp["Server won"], diff) {
		return "The Server won.", serverMove
	} else {
		return "Ooops, sth went wrong!", serverMove
	}
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /game", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "returns result of the game")
	})

	mux.HandleFunc("POST /game", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "post a new game")
	})

	mux.HandleFunc("POST /game/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintln(w, "post a new game with the move", moves[id])
		msg, servermove := generateMoveAndEvaluate(moves[id])
		fmt.Fprintln(w, "computer chose", servermove)
		fmt.Fprintln(w, msg)
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}
