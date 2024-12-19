package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/oxelf/pixel-party/internal/storage"
)

type Server struct {
	*storage.RedisConnection
}

func New() *Server {
	redis := storage.StartRedis(nil)
	return &Server{redis}
}

func (s *Server) Start() error {
	http.HandleFunc("/ws", handleWebSocket)
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/canvas", s.handleCanvasRequest)
	log.Println("Server started on :8080")
	return http.ListenAndServe(":8080", nil)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"Status": "Everythings looking fine",
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON", 500)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// WebSocket handling logic
}
