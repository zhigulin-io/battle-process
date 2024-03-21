package server

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type Game interface {
	GetState() any

	StartActivation(playerID int, unitID int) error
	StopActivation(playerID int) error
}

type Server struct {
	playerKeys []*uuid.UUID
	game       Game
}

func NewServer(game Game) *Server {
	return &Server{
		playerKeys: make([]*uuid.UUID, 2),
		game:       game,
	}
}

func (s *Server) ConnectAsPlayer(w http.ResponseWriter, _ *http.Request) {
	for i := range s.playerKeys {
		if s.playerKeys[i] == nil {
			response := struct {
				PlayerID int       `json:"playerID"`
				Key      uuid.UUID `json:"key"`
			}{
				PlayerID: i,
				Key:      uuid.New(),
			}
			writeSuccessResponse(w, response)
			s.playerKeys[i] = &response.Key
			return
		}
	}
	writeErrorResponse(w, http.StatusBadRequest, errors.New("no available slots"), "all players connected already")
}

func (s *Server) StartActivation(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("unitID") {
		writeErrorResponse(w, http.StatusBadRequest, errors.New("invalid request"), "query parameter unitID is required")
		return
	}

	unitID, err := strconv.Atoi(r.URL.Query().Get("unitID"))
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err, "cannot parse unitID to integer")
		return
	}

	playerID, err := s.detectPlayerID(r)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err, "cannot detect player id")
		return
	}

	err = s.game.StartActivation(playerID, unitID)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err, "cannot activate this unit")
		return
	}

	writeSuccessResponse(w, struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	})
}

func (s *Server) StopActivation(w http.ResponseWriter, r *http.Request) {
	playerID, err := s.detectPlayerID(r)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err, "cannot detect player id")
		return
	}

	err = s.game.StopActivation(playerID)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err, "cannot stop activation")
		return
	}

	writeSuccessResponse(w, struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	})
}

func (s *Server) GetGameState(w http.ResponseWriter, _ *http.Request) {
	writeSuccessResponse(w, s.game.GetState())
}
