package main

import (
	"context"
	"fmt"
	"github.com/rubenwo/SmartEnergyTable/Server/internal/pkg/session"
	"github.com/rubenwo/SmartEnergyTable/Server/internal/protocol"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type Server struct {
	SessionManager *session.Manager
}

func (s *Server) CreateSession(ctx context.Context, e *protocol.Create) (*protocol.Session, error) {
	id := s.SessionManager.CreateSession(e.User)
	return &protocol.Session{SessionId: id}, nil
}
func (s *Server) DestroySession(ctx context.Context, se *protocol.Session) (*protocol.Status, error) {
	ok, err := s.SessionManager.DestroySession(se.SessionId, se.User)
	if err != nil {
		log.Println(err)
		return &protocol.Status{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}, err
	}
	if !ok {

	}
	return &protocol.Status{}, nil
}
func (s *Server) StopSession(ctx context.Context, se *protocol.Session) (*protocol.Status, error) {
	return &protocol.Status{}, nil
}
func (s *Server) SaveSession(ctx context.Context, se *protocol.Session) (*protocol.Status, error) {
	_ = s.SessionManager.SaveSession(se.SessionId, se.User)
	return &protocol.Status{}, nil
}
func (s *Server) JoinSession(ctx context.Context, se *protocol.Session) (*protocol.Status, error) {
	_ = s.SessionManager.JoinSession(se.SessionId, se.User)
	return &protocol.Status{}, nil
}
func (s *Server) UpdateEntities(ctx context.Context, u *protocol.Update) (*protocol.Status, error) {
	return &protocol.Status{}, nil
}
func (s *Server) Refresh(ctx context.Context, se *protocol.Session) (*protocol.Update, error) {
	return &protocol.Update{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatal(err)
	}

	s := Server{
		SessionManager: session.NewManager(),
	}

	grpcServer := grpc.NewServer()
	protocol.RegisterSyncServiceServer(grpcServer, &s)
	log.Println("grpc server started listening on:", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
