package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"net"
	"rps2/engine/engine"
	"slices"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type GameEngineServer struct {
	engine.UnimplementedGameEngineServer
}

// len of client input minus len of server hardcoded who wins
var userMinusComp = map[string][]int{
	"Client won": {1, 2, -3},
	"Server won": {-1, -2, 3},
}

func generateMoveAndEvaluate(clientMove string) string {
	validInputs := []string{"rock", "paper", "scissor"}
	randNumber := rand.IntN(len(validInputs))
	serverMove := validInputs[randNumber]
	log.Printf("Move from server is %v", serverMove)
	diff := len(clientMove) - len(serverMove)

	if diff == 0 {
		return "Server and client tied"
	} else if slices.Contains(userMinusComp["Client won"], diff) {
		return "The Client won."
	} else if slices.Contains(userMinusComp["Server won"], diff) {
		return "The Server won."
	} else {
		return "Ooops, sth went wrong!"
	}

}

func (s *GameEngineServer) MakeMove(ctx context.Context, in *engine.MoveRequest) (*engine.MoveReply, error) {
	log.Printf("Move from client is %v", in.GetMove())
	return &engine.MoveReply{Winner: generateMoveAndEvaluate(in.GetMove())}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	engine.RegisterGameEngineServer(s, &GameEngineServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
