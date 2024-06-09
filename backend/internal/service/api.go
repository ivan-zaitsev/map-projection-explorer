package service

import (
	"map-projection-explorer-backend/internal/domain/dto"
)

type CrsService interface {
	FindCoordinateReferenceSystem(code int) (*dto.CrsRecord, *dto.ServiceError)
	FindAllCoordinateReferenceSystems(search string, pageCursor *int, pageSize int) (*dto.Page[dto.CrsRecord], *dto.ServiceError)
}
