package main

import (
	"context"
	"flag"
	"log"
	"math/rand/v2"
	"rps2/engine/engine"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func generateMove() string {
	validInputs := []string{"rock", "paper", "scissor"}
	randNumber := rand.IntN(len(validInputs))
	return validInputs[randNumber]
}

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	move = flag.String("move", generateMove(), "Move")
)

func main() {
	// set up connection to the server
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := engine.NewGameEngineClient(conn)

	// contact the server and print out its response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.MakeMove(ctx, &engine.MoveRequest{Move: *move})
	if err != nil {
		log.Fatalf("could not make move: %v", err)
	}
	log.Printf(r.GetWinner())
}
