package service

import "map-projection-explorer-backend/internal/domain/dto"

type CrsService interface {
	FindCoordinateReferenceSystem(code int) (*dto.CrsRecord, *dto.ServiceError)
	FindAllCoordinateReferenceSystems(cursorCode *int, pageSize int) ([]*dto.CrsRecord, *dto.ServiceError)
}
