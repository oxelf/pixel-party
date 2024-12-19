package server

import "net/http"

func (s *Server) handleCanvasRequest(w http.ResponseWriter, r *http.Request) {
	pixelBinary, err := s.GetCanvas("0")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)

	w.Write(pixelBinary)
}
