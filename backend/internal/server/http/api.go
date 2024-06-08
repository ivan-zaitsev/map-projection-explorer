package http_server

import "net/http"

type CrsHttpServer interface {
	FindCoordinateReferenceSystem(w http.ResponseWriter, r *http.Request)
	FindAllCoordinateReferenceSystems(w http.ResponseWriter, r *http.Request)
}

func RegisterServerHandlers(mux *http.ServeMux, server CrsHttpServer) {
	mux.HandleFunc("GET /api/v1/coordinate-reference-systems/{code}", server.FindCoordinateReferenceSystem)
	mux.HandleFunc("GET /api/v1/coordinate-reference-systems", server.FindAllCoordinateReferenceSystems)
}
