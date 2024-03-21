package server

import (
	"errors"
	"net/http"
)

func (s *Server) detectPlayerID(r *http.Request) (int, error) {
	playerKey := r.Header.Get("Player-Key")
	if playerKey == "" {
		return 0, errors.New("header Player-Key is required")
	}

	for i := range s.playerKeys {
		if s.playerKeys[i] != nil && playerKey == (*s.playerKeys[i]).String() {
			return i, nil
		}
	}

	return 0, errors.New("cannot find player by Player-Key")
}
