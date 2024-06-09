package http_server

import (
	"encoding/json"
	"log"
	"map-projection-explorer-backend/internal/domain/dto"
	"map-projection-explorer-backend/internal/service"
	"net/http"
	"strconv"
)

type crsServer struct {
	crsService service.CrsService
}

func NewServer(crsService service.CrsService) CrsHttpServer {
	return &crsServer{crsService: crsService}
}

func (c *crsServer) FindCoordinateReferenceSystem(w http.ResponseWriter, r *http.Request) {
	code, err := strconv.Atoi(r.PathValue("code"))
	if err != nil {
		serviceError := &dto.ServiceError{Code: dto.ErrorInvalidRequest, Message: "Request path value 'code' invalid"}
		writeResponse(w, serviceError, http.StatusBadRequest)
		return
	}

	crsRecord, serviceErr := c.crsService.FindCoordinateReferenceSystem(code)
	if serviceErr != nil {
		writeResponse(w, serviceErr, serviceErr.Code.Value)
		return
	}

	writeResponse(w, crsRecord, http.StatusOK)
}

func (c *crsServer) FindAllCoordinateReferenceSystems(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	search := query.Get("search")

	pageCursor, err := extractPageCursor(query.Get("pageCursor"))
	if err != nil {
		serviceError := &dto.ServiceError{Code: dto.ErrorInvalidRequest, Message: "Request parameter 'pageCursor' invalid"}
		writeResponse(w, serviceError, http.StatusBadRequest)
		return
	}

	pageSize, err := strconv.Atoi(query.Get("pageSize"))
	if err != nil {
		serviceError := &dto.ServiceError{Code: dto.ErrorInvalidRequest, Message: "Request parameter 'pageSize' invalid"}
		writeResponse(w, serviceError, http.StatusBadRequest)
		return
	}

	crsRecords, serviceErr := c.crsService.FindAllCoordinateReferenceSystems(search, pageCursor, pageSize)
	if serviceErr != nil {
		writeResponse(w, serviceErr, serviceErr.Code.Value)
		return
	}

	writeResponse(w, crsRecords, http.StatusOK)
}

func extractPageCursor(pageCursor string) (*int, error) {
	if pageCursor == "" {
		return nil, nil
	}
	val, err := strconv.Atoi(pageCursor)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

func writeResponse(w http.ResponseWriter, response any, status int) {
	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error serializing response", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(data)
	if err != nil {
		log.Println("Error writing response:", err)
	}
}
