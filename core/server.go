package core

import (
	"fmt"
	"quizmo/utils"

	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Server struct {
	Config  *utils.Config
	Echo    *echo.Echo
	Session neo4j.Session
}

func NewServer(config *utils.Config, db *utils.Database) *Server {

	// create session for neo4j
	session := db.Client.NewSession(neo4j.SessionConfig{
		DatabaseName: config.DatabaseName,
	})
	defer session.Close()

	server := &Server{
		Config:  config,
		Echo:    echo.New(),
		Session: session,
	}

	return server
}

func (s *Server) Start() {
	s.Echo.Logger.Fatal(s.Echo.Start(fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port)))
}
